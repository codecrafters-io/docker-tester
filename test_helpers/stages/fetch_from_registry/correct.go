package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"

	"github.com/mholt/archiver/v3"
)

func main() {
	if os.Args[1] != "run" {
		fmt.Printf("Expected 'run <image> <command>' as the command! args: %q\n", os.Args[1:])
		os.Exit(1)
	}

	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		fmt.Printf("TempDir Error: %v", err)
		os.Exit(1)
	}

	if err = downloadImageToPath(os.Args[2], tempDir); err != nil {
		fmt.Printf("Download Error: %v", err)
		os.Exit(1)
	}

	forkAttr := syscall.ProcAttr{
		Env: os.Environ(),
		Sys: &syscall.SysProcAttr{
			Chroot:     tempDir,
			Cloneflags: syscall.CLONE_NEWPID,
		},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	}

	pid, err := syscall.ForkExec(os.Args[3], os.Args[3:], &forkAttr)
	if err != nil {
		fmt.Printf("Fork Error: %v", err)
		os.Exit(1)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("FindProcess Error: %v", err)
		os.Exit(1)
	}

	state, err := process.Wait()
	if err != nil {
		fmt.Printf("ProcessWait Error: %v", err)
		os.Exit(1)
	}

	os.Exit(state.ExitCode())
}

func downloadImageToPath(image string, path string) error {
	url := fmt.Sprintf("http://localhost:5000/v2/%s/manifests/latest", image)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Expected 200, got: %d", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	// TODO: Cleanup JSON parsing
	for _, layer := range result["fsLayers"].([]interface{}) {
		digest := layer.(map[string]interface{})["blobSum"].(string)
		if err = downloadLayerToPath(image, digest, path); err != nil {
			return err
		}
	}

	return nil
}

func downloadLayerToPath(image string, digest string, path string) error {
	url := fmt.Sprintf("http://localhost:5000/v2/%s/blobs/%s", image, digest)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Expected 200, got: %d", resp.StatusCode)
	}

	tmpFile, err := ioutil.TempFile("", "*.tar.gz")
	if err != nil {
		return err
	}

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return err
	}

	return archiver.Unarchive(tmpFile.Name(), path)
}
