package wiki

import (
	"errors"
	"path/filepath"
)

const DefaultExtension = ".md"

type Wiki struct {
	// path is absolute path of wiki root in the filesystem
	path string

	// defaultExtension is a default markup extension used in this wiki
	// for now this is always .md
	defaultExtension string
}

func (w *Wiki) DefaultExtension() string {
	return w.defaultExtension
}

func New(path string) (*Wiki, error) {
	if !filepath.IsAbs(path) {
		return nil, errors.New("wiki_path must be absolute")
	}
	return &Wiki{
		path:             path,
		defaultExtension: DefaultExtension,
	}, nil
}
