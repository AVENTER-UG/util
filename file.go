package util

import (
	"io"
	"os"
)

// CopyFile will copy the file to a other destination
func CopyFile(from, to string) error {
	r, err := os.Open(from)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(to)
	if err != nil {
		return err
	}
	defer w.Close()

	// do the actual work
	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	return err
}

// MoveFile will move the file to a other destination
func MoveFile(from, to string) error {
	err := CopyFile(from, to)
	err = os.Remove(from)

	return err
}

