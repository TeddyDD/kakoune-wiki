package common

import (
	"fmt"
	"path/filepath"
	"strings"
)

// AppendExtension to given path. If extension already present, return
// unmodified path
func AppendExtension(path, ext string) string {
	if strings.HasSuffix(path, ext) {
		return path
	}

	return fmt.Sprintf("%s%s", path, ext)
}

// TrimExtension to given path. If extension already present, return
// unmodified path
func TrimExtension(path, ext string) string {
	if !strings.HasSuffix(path, ext) {
		return path
	}

	return strings.TrimSuffix(path, ext)
}

// relativePath returns relative path for linking dest file from
// src file
func RelativePath(dest, src string) (string, error) {
	srcDir := filepath.Dir(src)
	return filepath.Rel(srcDir, dest)
}
