package md

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	FileName string
	BodyRaw []byte
	BodyHTML []byte
}

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func Posts(path string) (map[string]Post, error) {
	source := os.DirFS(path)
	dirEntries, err := fs.Glob(source, "*.md")
	if err != nil {
		return nil, err
	} 

	posts := make(map[string]Post)

	for _, p := range dirEntries {
		mdPath := fmt.Sprintf("%s/%s", path, p)
		body, _ := os.ReadFile(mdPath)
		post := Post {
			FileName: p,
			BodyHTML: mdToHTML(body),
			BodyRaw: body,
		}

		posts[p] = post
	}

	return posts, nil
}
