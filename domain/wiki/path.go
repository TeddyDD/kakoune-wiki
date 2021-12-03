package wiki

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/TeddyDD/kakoune-wiki/domain/common"
)

// AbsolutePath returns absolute path assuming path is
// in wiki directory
func (w *Wiki) AbsolutePath(path string) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}

	path = filepath.Join("/", path)
	path = filepath.Clean(path)
	path = filepath.Join(w.path, path)
	return filepath.Clean(path)
}

// RelativeToWiki makes path relative to wiki directory
func (w *Wiki) RelativeToWiki(path string) (string, error) {
	if !filepath.IsAbs(path) {
		return path, nil
	}
	path, err := filepath.Rel(w.path, path)
	if err != nil {
		return "", fmt.Errorf("%s not in wiki directory: %w", path, err)
	}
	return path, nil
}

// FileInWiki checks is given path (usually buffile) is in wiki.
// TODO: move to Kakoune
func (w *Wiki) FileInWiki(path string) bool {
	path = filepath.Clean(path)
	return strings.HasPrefix(path, w.path)
}

// AddresToMarkdown converts MediaWiki style addres to markdown.
// inFile is a absolute path to file in wiki that contains the addres
// mediaWikiAddres is relative to wiki root
func (w *Wiki) AddresToMarkdown(mediaWikiAddres, inFile string) string {
	md := common.AppendExtension(mediaWikiAddres, w.defaultExtension)
	// mediaWikiAddres is always relative to root of wiki
	// but Markdown requires relative path between two files
	md, _ = common.RelativePath(md, inFile)
	return md
}

func (w *Wiki) AddresToMediaWiki(markdownAddres, inFile string) string {
	mediaWiki := filepath.Join(filepath.Dir(inFile), markdownAddres)
	mediaWiki = common.TrimExtension(mediaWiki, w.defaultExtension)

	return mediaWiki
}
