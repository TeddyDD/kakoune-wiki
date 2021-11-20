package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func relativeCmd(c *config, dest, src string) error {
	err := os.Chdir(c.WikiPath)
	if err != nil {
		return err
	}
	rel, err := relativePath(dest, src)
	if err != nil {
		return err
	}

	fmt.Println(rel)
	return nil
}

// relativePath returns relative path for linking dest file from
// src file
func relativePath(dest, src string) (string, error) {
	srcDir := filepath.Dir(src)
	return filepath.Rel(srcDir, dest)
}
