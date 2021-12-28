package wiki

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func (w *Wiki) Files() ([]string, error) {
	// TODO: add cache?
	files := make([]string, 0, 100)
	err := filepath.WalkDir(w.path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files, err
}

func FilterMarkdown(files []string) []string {
	res := make([]string, 0, len(files))
	for i := range files {
		base := filepath.Base(files[i])
		ext := strings.ToLower(filepath.Ext(base))
		if ext == DefaultExtension {
			res = append(res, files[i])
		}
	}
	return res
}

func FilterPrefixNoCase(files []string, prefix string) []string {
	res := make([]string, 0, len(files))
	prefix = strings.ToLower(prefix)
	for i := range files {
		base := strings.ToLower(filepath.Base(files[i]))
		if strings.HasPrefix(base, prefix) {
			res = append(res, files[i])
		}
	}
	return res
}
