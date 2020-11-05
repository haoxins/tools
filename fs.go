package tools

import (
	"io"
	"os"
	"path/filepath"
)

// CopyFile copy file
func CopyFile(from, to string) error {
	src, err := os.Open(from)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

// Getwd ...
func Getwd() string {
	dir, err := os.Getwd()
	PanicError(err)
	return dir
}

// Resolve ...
func Resolve(from string, to string) string {
	if filepath.IsAbs(to) {
		return to
	}

	if filepath.IsAbs(from) {
		return filepath.Join(from, to)
	}

	abs, err := filepath.Abs(from)

	if err != nil {
		abs = "/"
	}

	return filepath.Join(abs, to)
}
