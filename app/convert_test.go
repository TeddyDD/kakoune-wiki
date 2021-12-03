package app

import (
	"testing"

	"github.com/TeddyDD/kakoune-wiki/domain/kakoune"
	"github.com/TeddyDD/kakoune-wiki/domain/wiki"
	"github.com/stretchr/testify/require"
)

func Test_app_ConvertMediaWikiLinkToMarkdown(t *testing.T) {
	testWiki, _ := wiki.New("/wiki")

	tests := []struct {
		name      string
		buffile   string
		inputLink string
		want      string
		assertion require.ErrorAssertionFunc
	}{
		{
			name:      "simple",
			buffile:   "/wiki/index.md",
			inputLink: "[[foo]]",
			want:      "[foo](foo.md)",
			assertion: require.NoError,
		},
		{
			name:      "simple - custom alt text",
			buffile:   "/wiki/index.md",
			inputLink: "[[Hello|foo]]",
			want:      "[Hello](foo.md)",
			assertion: require.NoError,
		},
		{
			name:      "linked from root to subdirectory",
			buffile:   "/wiki/index.md",
			inputLink: "[[foo/bar]]",
			want:      "[foo/bar](foo/bar.md)",
			assertion: require.NoError,
		},
		{
			name:      "linked form subdirectory to root",
			buffile:   "/wiki/foo/bar.md",
			inputLink: "[[index]]",
			want:      "[index](../index.md)",
			assertion: require.NoError,
		},
		{
			name:      "linked form subdirectory to root - custom alt",
			buffile:   "/wiki/foo/bar.md",
			inputLink: "[[Hello|index]]",
			want:      "[Hello](../index.md)",
			assertion: require.NoError,
		},
		{
			name:      "subdirectory sibling",
			buffile:   "/wiki/foo/foo.md",
			inputLink: "[[bar/bar]]",
			want:      "[bar/bar](../bar/bar.md)",
			assertion: require.NoError,
		},
		{
			name:      "empty buffile",
			buffile:   "",
			inputLink: "[[foo]]",
			want:      "",
			assertion: require.Error,
		},
		{
			name:      "buffile outisde of wiki path",
			buffile:   "/app/foo.md",
			inputLink: "[[foo]]",
			want:      "",
			assertion: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := app{
				config: &kakoune.Config{
					Buffile: tt.buffile,
				},
				wiki: testWiki,
			}
			got, err := a.ConvertMediaWikiLinkToMarkdown(tt.inputLink)
			tt.assertion(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
