package main

import (
    "bytes"
    "fmt"
    "github.com/mattn/go-isatty"
)


type FakeFile struct {
	b bytes.Buffer
}

// NewFakeFile creates a FakeFile
func NewFakeFile() *FakeFile {
	return &FakeFile{}
}

// Fd returns the file descriptor
func (f *FakeFile) Fd() uintptr {
	return uintptr(0)
}

func main() {
	f := NewFakeFile()
	fmt.Println("IsTerminal: ", isatty.IsTerminal(f.Fd()))
}
