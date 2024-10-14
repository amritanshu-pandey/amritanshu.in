package backend

import (
	"fmt"
	"net/http"

	"amritanshu.in/goblog/md"
	"amritanshu.in/goblog/views"
	"github.com/a-h/templ"
)

func RunServer(markdownPath string, serverPort int, serverBindAddr string) error {
	markdownPosts, err := md.ActivePosts(markdownPath)
	if err != nil {
		return err
	}
	sortedTitles, err := md.SortedPostsByDate(markdownPath)
	if err != nil {
		return err
	}

	http.Handle("/", templ.Handler(views.Index(markdownPosts, sortedTitles)))
	http.HandleFunc("/article/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		views.Article(markdownPosts[slug]).Render(r.Context(), w)
	})
	http.ListenAndServe(fmt.Sprintf("%s:%d", serverBindAddr, serverPort), nil)
	return nil
}
