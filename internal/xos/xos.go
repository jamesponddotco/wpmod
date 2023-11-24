// Package xos contains functions and utilities to extend the os module from
// Go's standard library.
package xos

import (
	"fmt"
	"os"
	"path/filepath"
)

// Rchown changes the owner and group of the given path recursively.
func Rchown(basePath string, uid, gid int) error {
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking %q: %w", path, err)
		}

		if err := os.Chown(path, uid, gid); err != nil {
			return fmt.Errorf("error changing owner and group for %q: %w", path, err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Rchmod changes the mode of the given path recursively. It supports different
// modes for files and directories.
func Rchmod(basePath string, fileMode, directoryMode os.FileMode) error {
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking %q: %w", path, err)
		}

		if info.IsDir() {
			if err := os.Chmod(path, directoryMode); err != nil {
				return fmt.Errorf("error changing mode for %q: %w", path, err)
			}
		} else {
			if err := os.Chmod(path, fileMode); err != nil {
				return fmt.Errorf("error changing mode for %q: %w", path, err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
