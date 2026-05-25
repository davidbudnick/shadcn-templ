// Command example is a minimal, self-contained app that consumes the
// shadcn-templ library — the smallest realistic starting point for your own
// project. Run it with:
//
//	go run ./example
//
// then open http://localhost:8081.
package main

import (
	"log"
	"net/http"

	shadcntempl "github.com/davidbudnick/shadcn-templ"
)

func main() {
	mux := http.NewServeMux()

	// Serve the stylesheet that ships embedded in the library.
	mux.HandleFunc("GET /static/css/styles.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		_, _ = w.Write(shadcntempl.CSS)
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = dashboardPage().Render(r.Context(), w)
	})

	const addr = ":8081"
	log.Printf("example app listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalln(err)
	}
}
