package util

import (
	"io"
	"os"
	"crypto/sha256"
	"encoding/hex"
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

// GetFileHash return a sha256 hashvalue of a file
func GetFileHash(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := sha256.New()

	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
