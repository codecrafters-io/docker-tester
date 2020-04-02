package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"os"
	"os/exec"
	"strings"
	"syscall"
)

var REGISTRY_BASE_URL = "https://registry.hub.docker.com"

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

	var imageName = os.Args[2]
	var command = os.Args[3]
	var commandArgs = os.Args[3:]

	if !strings.Contains(imageName, "/") {
		imageName = "library/" + imageName
	}

	token, err := fetchToken(imageName)
	if err != nil {
		fmt.Printf("Token fetch failure: %v", err)
		os.Exit(1)
	}

	if err = downloadImageToPath(imageName, token, tempDir); err != nil {
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

	pid, err := syscall.ForkExec(command, commandArgs, &forkAttr)
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

func fetchToken(image string) (string, error) {
	url := fmt.Sprintf(REGISTRY_BASE_URL+"/v2/%s/manifests/latest", image)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 401 {
		return "", fmt.Errorf("Expected 401, got: %d", resp.StatusCode)
	}

	var headerValue = resp.Header.Get("Www-Authenticate")
	if !strings.HasPrefix(headerValue, "Bearer realm") {
		return "", fmt.Errorf("Expected valid auth header, got: %d", headerValue)
	}

	var headerSplit = strings.SplitN(headerValue, "\"", -1)
	var realm, service, scope = headerSplit[1], headerSplit[3], headerSplit[5]
	var authUrl = fmt.Sprintf("%s?service=%s&scope=%s", realm, service, scope)

	resp, err = http.Get(authUrl)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Expected 200, got: %d", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(bytes, &result); err != nil {
		return "", err
	}
	return result["token"].(string), nil
}

func downloadImageToPath(image string, token string, path string) error {
	url := fmt.Sprintf(REGISTRY_BASE_URL+"/v2/%s/manifests/latest", image)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		fmt.Println(string(bytes))
		return fmt.Errorf("Expected 200, got: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err = json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	// TODO: Cleanup JSON parsing
	for _, layer := range result["fsLayers"].([]interface{}) {
		digest := layer.(map[string]interface{})["blobSum"].(string)
		if err = downloadLayerToPath(client, token, image, digest, path); err != nil {
			return err
		}
	}

	return nil
}

func downloadLayerToPath(client *http.Client, token string, image string, digest string, path string) error {
	url := fmt.Sprintf(REGISTRY_BASE_URL+"/v2/%s/blobs/%s", image, digest)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
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

	cmd := exec.Command("tar", "-xf", tmpFile.Name(), "-C", path)
	return cmd.Run()
}
