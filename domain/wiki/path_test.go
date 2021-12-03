package wiki

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const testWikiPath = "/wiki"

func testWiki() *Wiki {
	return &Wiki{
		path:             testWikiPath,
		defaultExtension: DefaultExtension,
	}
}

func TestWiki_AbsolutePath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "simple",
			path: "foo",
			want: "/wiki/foo",
		},
		{
			name: "simple absolute",
			path: "/wiki/foo",
			want: "/wiki/foo",
		},
		{
			name: "escape attempt",
			path: "../../foo",
			want: "/wiki/foo",
		},
		{
			name: "garbage",
			path: "././foo",
			want: "/wiki/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := testWiki()
			require.Equal(t, tt.want, w.AbsolutePath(tt.path))
		})
	}
}

func TestWiki_FileInWiki(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "wiki is in wiki",
			path: testWikiPath,
			want: true,
		},
		{
			name: "simple",
			path: filepath.Join(testWikiPath, "foo"),
			want: true,
		},
		{
			name: "not in wiki",
			path: "/usr/wiki",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := testWiki()
			require.Equal(t, tt.want, w.FileInWiki(tt.path))
		})
	}
}

func TestWiki_AddresToMarkdown(t *testing.T) {
	type args struct {
		mediaWikiAddres string
		inFile          string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "root dir - simple",
			args: args{
				mediaWikiAddres: "foo",
				inFile:          "index.md",
			},
			want: "foo.md",
		},
		{
			name: "subdirectory linked from root",
			args: args{
				mediaWikiAddres: "foo/bar",
				inFile:          "index.md",
			},
			want: "foo/bar.md",
		},
		{
			name: "root linked from subdirectory",
			args: args{
				mediaWikiAddres: "index",
				inFile:          "foo/bar.md",
			},
			want: "../index.md",
		},
		{
			name: "sibling subdirectory",
			args: args{
				mediaWikiAddres: "foo/bar",
				inFile:          "asdf/bar.md",
			},
			want: "../foo/bar.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := testWiki()
			require.Equal(t, tt.want, w.AddresToMarkdown(tt.args.mediaWikiAddres, tt.args.inFile))
		})
	}
}

func TestWiki_AddresToMediaWiki(t *testing.T) {
	type args struct {
		markdownAddres string
		inFile         string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sibling in root dir",
			args: args{
				markdownAddres: "index.md",
				inFile:         "foo.md",
			},
			want: "index",
		},
		{
			name: "subdirectory linked from root",
			args: args{
				markdownAddres: "foo/bar.md",
				inFile:         "index.md",
			},
			want: "foo/bar",
		},
		{
			name: "root linked form subdirectory",
			args: args{
				markdownAddres: "../index.md",
				inFile:         "foo/bar.md",
			},
			want: "index",
		},
		{
			name: "subdirectory sibling",
			args: args{
				markdownAddres: "bar.md",
				inFile:         "foo/baz.md",
			},
			want: "foo/bar",
		},
		{
			name: "different subdirectory sibling",
			args: args{
				markdownAddres: "../asdf/bar.md",
				inFile:         "foo/baz.md",
			},
			want: "asdf/bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := testWiki()
			require.Equal(t, tt.want, w.AddresToMediaWiki(tt.args.markdownAddres, tt.args.inFile))
		})
	}
}
