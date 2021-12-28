package app

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/TeddyDD/kakoune-wiki/domain/kakoune"
	"github.com/TeddyDD/kakoune-wiki/domain/wiki"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealList(t *testing.T) {
	abs, err := filepath.Abs("../testdata")
	require.NoError(t, err)

	w, err := wiki.New(abs)
	require.NoError(t, err)
	files, err := w.Files()
	require.NoError(t, err)
	t.Logf("%+v\n", files)
}

var testFiles = func() ([]string, error) {
	return []string{
		// Do not test files outisde of the wiki path since
		// they won't be returned in app
		"/wiki/index.md",
		"/wiki/testFile.md",
		"/wiki/foo.pdf",
		"/wiki/foo/foo.md",
		"/wiki/foo/bar.md",
		"/wiki/abc/abc.md",
		"/wiki/abc/bar.md",
		"/wiki/abc/123.md",
	}, nil
}

func Test_app_CompleteMediaWiki(t *testing.T) {
	type fields struct {
		config *kakoune.Config
		wiki   *wiki.Wiki
		files  func() ([]string, error)
	}

	wiki, err := wiki.New("/wiki")
	require.NoError(t, err)

	defaultFields := fields{
		wiki:  wiki,
		files: testFiles,
	}

	type args struct {
		prefix string
	}
	tests := []struct {
		name      string
		fields    *fields
		args      args
		want      kakoune.Completions
		assertion require.ErrorAssertionFunc
	}{
		{
			name: "index",
			args: args{
				prefix: "index",
			},
			want: []kakoune.CompletionEntry{
				kakoune.NewCompletionEntry("index", "index.md"),
			},
			assertion: require.NoError,
		},
		{
			name: "index short",
			args: args{
				prefix: "ind",
			},
			want: []kakoune.CompletionEntry{
				kakoune.NewCompletionEntry("index", "index.md"),
			},
			assertion: require.NoError,
		},
		{
			name: "nested",
			args: args{
				prefix: "abc",
			},
			want: []kakoune.CompletionEntry{
				kakoune.NewCompletionEntry("abc/abc", "abc/abc.md"),
				kakoune.NewCompletionEntry("abc/bar", "abc/bar.md"),
				kakoune.NewCompletionEntry("abc/123", "abc/123.md"),
			},
			assertion: require.NoError,
		},
		{
			name: "foo",
			args: args{
				prefix: "foo",
			},
			want: []kakoune.CompletionEntry{
				kakoune.NewCompletionEntry("foo/foo", "foo/foo.md"),
				kakoune.NewCompletionEntry("foo/bar", "foo/bar.md"),
			},
			assertion: require.NoError,
		},
		{
			name: "multiple",
			args: args{
				prefix: "bar",
			},
			want: []kakoune.CompletionEntry{
				kakoune.NewCompletionEntry("foo/bar", "foo/bar.md"),
				kakoune.NewCompletionEntry("abc/bar", "abc/bar.md"),
			},
			assertion: require.NoError,
		},
		{
			name: "empty",
			args: args{
				prefix: "zzz",
			},
			want:      []kakoune.CompletionEntry{},
			assertion: require.NoError,
		},
		{
			name: "error bubble up",
			fields: &fields{
				wiki: wiki,
				files: func() ([]string, error) {
					return nil, errors.New("asdf")
				},
			},
			args:      args{},
			assertion: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f *fields

			if tt.fields != nil {
				f = tt.fields
			} else {
				f = &defaultFields
			}
			a := app{
				config: f.config,
				wiki:   f.wiki,
				files:  f.files,
			}

			got, err := a.CompleteMediaWiki(tt.args.prefix)
			tt.assertion(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_app_CompleteMarkdown(t *testing.T) {
	type fields struct {
		config *kakoune.Config
		wiki   *wiki.Wiki
		files  func() ([]string, error)
	}

	wiki, err := wiki.New("/wiki")
	require.NoError(t, err)

	defaultFields := fields{
		wiki:  wiki,
		files: testFiles,
	}
	mkfields := func(buffile string) *fields {
		return &fields{
			config: &kakoune.Config{
				Buffile: buffile,
			},
			wiki:  wiki,
			files: testFiles,
		}
	}

	type args struct {
		prefix string
	}
	tests := []struct {
		name      string
		fields    *fields
		args      args
		want      kakoune.Completions
		assertion require.ErrorAssertionFunc
	}{
		{
			name:   "sibling",
			fields: mkfields("/wiki/index.md"),
			args: args{
				prefix: "test",
			},
			want: []kakoune.CompletionEntry{
				kakoune.NewCompletionEntry("testFile.md", "testFile.md"),
			},
		},
		{
			name:   "child",
			fields: mkfields("/wiki/index.md"),
			args: args{
				prefix: "foo/foo",
			},
			want: []kakoune.CompletionEntry{
				kakoune.NewCompletionEntry("foo/foo.md", "foo/foo.md"),
			},
		},
		{
			name:   "parent",
			fields: mkfields("/wiki/abc/abc.md"),
			args: args{
				prefix: "test",
			},
			want: []kakoune.CompletionEntry{
				kakoune.NewCompletionEntry("../testFile.md", "testFile.md"),
			},
		},
		{
			name:   "subdir sibling",
			fields: mkfields("/wiki/abc/abc.md"),
			args: args{
				prefix: "foo/foo",
			},
			want: []kakoune.CompletionEntry{
				kakoune.NewCompletionEntry("../foo/foo.md", "foo/foo.md"),
			},
		},
		{
			name:   "bad buffile",
			fields: mkfields("/foo/abc/abc.md"),
			args: args{
				prefix: "foo/foo",
			},
			assertion: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f *fields
			if tt.fields != nil {
				f = tt.fields
			} else {
				f = &defaultFields
			}
			a := app{
				config: f.config,
				wiki:   f.wiki,
				files:  f.files,
			}

			if tt.assertion == nil {
				tt.assertion = require.NoError
			}

			got, err := a.CompleteMarkdown(tt.args.prefix)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
