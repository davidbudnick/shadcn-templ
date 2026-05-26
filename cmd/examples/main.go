// Command examples runs (and can statically export) the documentation site for
// the shadcn-templ component library. Every component has a page mirroring its
// shadcn/ui demo, and the full page-level examples (Authentication, Dashboard,
// Mail, ...) are reproduced as well.
//
// Run the dev server:
//
//	go run ./cmd/examples
//
// Export a static site for hosting (e.g. GitHub Pages):
//
//	go run ./cmd/examples -build ./dist -base /shadcn-templ
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/a-h/templ"
	shadcntempl "github.com/davidbudnick/shadcn-templ"
)

// BasePath is prefixed to every internal link. It is empty for the dev server
// and set to e.g. "/shadcn-templ" when exporting for a GitHub Pages project
// site served from https://<user>.github.io/shadcn-templ/.
var BasePath = ""

// NavItem is a single entry in a navigation list.
type NavItem struct {
	Slug  string
	Title string
}

// Example is one rendered demo within a component's page. Func is the name of
// the templ function that produces Component; its source is shown as the code
// sample on the page.
type Example struct {
	Name      string
	Func      string
	Component templ.Component
}

// Section is a component's documentation page.
type Section struct {
	Slug        string
	Title       string
	Description string
	Examples    []Example
}

// ExamplePage is a full-page example reproduced from the shadcn/ui "Examples".
type ExamplePage struct {
	Slug      string
	Title     string
	Component templ.Component
}

// examplePages is the registry of full-page examples, in sidebar order.
var examplePages = []ExamplePage{
	{Slug: "dashboard", Title: "Dashboard", Component: DashboardExample()},
	{Slug: "authentication", Title: "Authentication", Component: AuthenticationExample()},
	{Slug: "cards", Title: "Cards", Component: CardsExample()},
	{Slug: "tasks", Title: "Tasks", Component: TasksExample()},
	{Slug: "forms", Title: "Forms", Component: SettingsExample()},
	{Slug: "music", Title: "Music", Component: MusicExample()},
	{Slug: "mail", Title: "Mail", Component: MailExample()},
}

func componentNav() []NavItem {
	items := make([]NavItem, 0, len(sections))
	for _, s := range sections {
		items = append(items, NavItem{Slug: s.Slug, Title: s.Title})
	}
	return items
}

func examplesNav() []NavItem {
	items := make([]NavItem, 0, len(examplePages))
	for _, p := range examplePages {
		items = append(items, NavItem{Slug: p.Slug, Title: p.Title})
	}
	return items
}

// themeColor is a selectable accent theme and the HSL of its primary color
// (used for the swatch). Matches the presets in static/css/input.css.
type themeColor struct {
	Name  string // data-theme value ("" = default Zinc)
	Label string
	HSL   string
}

// themeColors are the visually-distinct options shown in the docs theme picker
// (Zinc default + accent colors). The softer base palettes — slate, stone,
// gray, neutral — are also defined in the CSS and can be set via
// data-theme="…", they're just omitted here because their swatches all read as
// the same dark dot.
var themeColors = []themeColor{
	{"", "Zinc", "240 5.9% 10%"},
	{"red", "Red", "0 72.2% 50.6%"},
	{"rose", "Rose", "346.8 77.2% 49.8%"},
	{"orange", "Orange", "24.6 95% 53.1%"},
	{"green", "Green", "142.1 76.2% 36.3%"},
	{"blue", "Blue", "221.2 83.2% 53.3%"},
	{"yellow", "Yellow", "47.9 95.8% 53.1%"},
	{"violet", "Violet", "262.1 83.3% 57.8%"},
}

// --- URL helpers (honor BasePath so the same templates work on the dev server
// and on a GitHub Pages sub-path) ---

func urlHome() templ.SafeURL              { return templ.SafeURL(BasePath + "/") }
func urlComponent(s string) templ.SafeURL { return templ.SafeURL(BasePath + "/components/" + s) }
func urlExample(s string) templ.SafeURL   { return templ.SafeURL(BasePath + "/examples/" + s) }
func urlCSS() string                      { return BasePath + "/static/css/styles.css" }

func itoa(n int) string { return strconv.Itoa(n) }

func main() {
	addr := flag.String("addr", ":8080", "address to listen on (dev server)")
	buildDir := flag.String("build", "", "if set, export a static site to this directory and exit")
	base := flag.String("base", "", "base path prefix for links (e.g. /shadcn-templ for GitHub Pages)")
	flag.Parse()

	BasePath = *base

	if *buildDir != "" {
		if err := buildStaticSite(*buildDir); err != nil {
			log.Fatalln("build:", err)
		}
		log.Printf("static site written to %s (base %q)", *buildDir, BasePath)
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /static/css/styles.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		_, _ = w.Write(shadcntempl.CSS)
	})

	cNav := componentNav()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		render(w, r, Home(cNav))
	})

	mux.HandleFunc("GET /components/{slug}", func(w http.ResponseWriter, r *http.Request) {
		for _, s := range sections {
			if s.Slug == r.PathValue("slug") {
				render(w, r, ComponentPage(cNav, s))
				return
			}
		}
		http.NotFound(w, r)
	})

	mux.HandleFunc("GET /examples/{slug}", func(w http.ResponseWriter, r *http.Request) {
		for _, p := range examplePages {
			if p.Slug == r.PathValue("slug") {
				render(w, r, exampleDocument(examplesNav(), p.Slug, p.Title, p.Component))
				return
			}
		}
		http.NotFound(w, r)
	})

	log.Printf("examples site listening on http://localhost%s", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatalln(err)
	}
}

func render(w http.ResponseWriter, r *http.Request, c templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// buildStaticSite renders every page to disk as a self-contained static site
// that can be served by any static host (GitHub Pages, Netlify, S3, ...).
func buildStaticSite(dir string) error {
	cNav := componentNav()
	eNav := examplesNav()

	// Stylesheet + GitHub Pages opt-out of Jekyll processing.
	if err := writeFile(filepath.Join(dir, "static", "css", "styles.css"), func(w io.Writer) error {
		_, err := w.Write(shadcntempl.CSS)
		return err
	}); err != nil {
		return err
	}
	if err := writeFile(filepath.Join(dir, ".nojekyll"), func(w io.Writer) error { return nil }); err != nil {
		return err
	}

	// Home.
	if err := writePage(filepath.Join(dir, "index.html"), Home(cNav)); err != nil {
		return err
	}
	// Component pages.
	for _, s := range sections {
		if err := writePage(filepath.Join(dir, "components", s.Slug, "index.html"), ComponentPage(cNav, s)); err != nil {
			return err
		}
	}
	// Full-page examples.
	for _, p := range examplePages {
		if err := writePage(filepath.Join(dir, "examples", p.Slug, "index.html"), exampleDocument(eNav, p.Slug, p.Title, p.Component)); err != nil {
			return err
		}
	}
	return nil
}

func writePage(path string, c templ.Component) error {
	return writeFile(path, func(w io.Writer) error {
		return c.Render(context.Background(), w)
	})
}

func writeFile(path string, write func(io.Writer) error) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return write(f)
}
