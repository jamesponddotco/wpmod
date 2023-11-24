// Package find provides functions to find files and directories.
package find

import (
	"fmt"
	"os"
	"path/filepath"
)

// MuPluginsDirectory traverses the given directory tree and returns the path to
// the first directory named "mu-plugins" that is found.
func MuPluginsDirectory(root string) (string, error) {
	var muPluginsPath string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking path %q: %w", path, err)
		}

		if info.IsDir() && info.Name() == "mu-plugins" {
			muPluginsPath = path

			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return muPluginsPath, nil
}
