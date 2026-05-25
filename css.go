package shadcntempl

import _ "embed"

// CSS is the compiled CSS used in this module.
//
// You need to serve this from a path and link to it in your <head> tag.
//
//	router.HandleFunc("GET /static/css/styles.css", func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "text/css")
//		w.Write(shadcntempl.CSS)
//	})
//
//go:embed static/css/output.css
var CSS []byte
