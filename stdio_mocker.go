package main

import (
	"io/ioutil"
	"os"
)

type IOMocker struct {
	originalStdout *os.File
	originalStderr *os.File
	originalStdin  *os.File
}

func NewStdIOMocker() *IOMocker {
	return &IOMocker{}
}

func (m *IOMocker) Start() {
	m.originalStdout = os.Stdout
	m.originalStdin = os.Stdin
	m.originalStdin = os.Stdin

	os.Stdout, _ = ioutil.TempFile("", "")
	os.Stdin, _ = ioutil.TempFile("", "")
	os.Stderr, _ = ioutil.TempFile("", "")
}

func (m *IOMocker) End() {
	m.originalStdout = os.Stdout
	m.originalStdin = os.Stdin
	m.originalStdin = os.Stdin

	os.Stdout = m.originalStdout
	os.Stdin = m.originalStdin
	os.Stderr = m.originalStderr
}
