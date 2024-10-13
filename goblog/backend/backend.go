package backend

import (
	"net/http"

	"amritanshu.in/goblog/md"
	"amritanshu.in/goblog/views"
	"github.com/a-h/templ"
)

func RunServer() error {
	posts, err := md.Posts()
	if err != nil {
		return err
	}
	http.Handle("/", templ.Handler(views.Index(posts)))
	http.Handle("/article/", templ.Handler(views.Index(posts)))
	http.HandleFunc("/article/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		views.Article(posts[slug]).Render(r.Context(), w)
	})
	http.ListenAndServe(":8080", nil)
	return nil
}