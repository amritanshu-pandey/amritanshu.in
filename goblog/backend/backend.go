package backend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"amritanshu.in/goblog/md"
	"amritanshu.in/goblog/views"
	"github.com/a-h/templ"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func initMeter() (*sdkmetric.MeterProvider, error) {
	exp, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exp)))
	otel.SetMeterProvider(mp)
	return mp, nil
}

func RunServer(markdownPath string, assetsDir string, serverPort int, serverBindAddr string) error {
	mux := http.NewServeMux()
	markdownPosts, err := md.ActivePosts(markdownPath)
	if err != nil {
		return err
	}
	sortedTitles, err := md.SortedPostsByDate(markdownPath)
	if err != nil {
		return err
	}

	mp, err := initMeter()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down meter provider: %v", err)
		}
	}()

	// Index Page
	http.Handle("/", otelhttp.NewHandler(otelhttp.WithRouteTag("/", templ.Handler(views.Index(markdownPosts, sortedTitles))), ""))

	// Static assets
	staticFs := http.FileServer(http.Dir(assetsDir))
	http.Handle("/assets/", otelhttp.NewHandler(otelhttp.WithRouteTag("/assets", http.StripPrefix("/assets/", staticFs)), ""))

	// Articles
	// http.HandleFunc("/article/{slug}", func(w http.ResponseWriter, r *http.Request) {
	// 	slug := r.PathValue("slug")
	// 	views.Article(markdownPosts[slug]).Render(r.Context(), w)
	// })

	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		handler:= otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	handleFunc("/article/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		views.Article(markdownPosts[slug]).Render(r.Context(), w)
	})

	http.ListenAndServe(fmt.Sprintf("%s:%d", serverBindAddr, serverPort), nil)
	return nil
}
