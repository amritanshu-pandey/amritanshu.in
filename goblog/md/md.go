package md

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

//go:embed *.md
var source embed.FS

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

func Posts() (map[string]Post, error) {
	dirEntries, err := source.ReadDir(".")
	if err != nil {
		return nil, err
	} 

	posts := make(map[string]Post)

	for _, p := range dirEntries {
		body, _ := source.ReadFile(p.Name())
		post := Post{
			FileName: p.Name(),
			BodyHTML: mdToHTML(body),
			BodyRaw: body,
		}

		posts[p.Name()] = post
	}

	return posts, nil
}
