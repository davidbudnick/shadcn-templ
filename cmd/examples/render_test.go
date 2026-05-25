package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

// TestSectionsRender renders every component's page, which in turn renders
// every demo and therefore exercises every UI component in the library.
func TestSectionsRender(t *testing.T) {
	cNav := componentNav()
	for _, s := range sections {
		t.Run(s.Slug, func(t *testing.T) {
			var sb strings.Builder
			if err := ComponentPage(cNav, s).Render(context.Background(), &sb); err != nil {
				t.Fatalf("render: %v", err)
			}
			out := sb.String()
			if !strings.Contains(out, s.Title) {
				t.Errorf("output missing component title %q", s.Title)
			}
			if len(out) < 500 {
				t.Errorf("suspiciously short output (%d bytes)", len(out))
			}
		})
	}
}

// TestExamplePagesRender renders every full-page example.
func TestExamplePagesRender(t *testing.T) {
	eNav := examplesNav()
	for _, p := range examplePages {
		t.Run(p.Slug, func(t *testing.T) {
			var sb strings.Builder
			if err := exampleDocument(eNav, p.Slug, p.Title, p.Component).Render(context.Background(), &sb); err != nil {
				t.Fatalf("render: %v", err)
			}
			out := sb.String()
			if !strings.Contains(out, "<!doctype html>") && !strings.Contains(out, "<!DOCTYPE html>") {
				t.Errorf("example %q did not render a full document", p.Slug)
			}
			if len(out) < 1000 {
				t.Errorf("suspiciously short output (%d bytes)", len(out))
			}
		})
	}
}

func TestHomeRenders(t *testing.T) {
	var sb strings.Builder
	if err := Home(componentNav()).Render(context.Background(), &sb); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(sb.String(), "shadcn-templ") {
		t.Error("home page missing brand name")
	}
}

// TestRegistryComplete guards against half-registered components and examples.
func TestRegistryComplete(t *testing.T) {
	seen := map[string]bool{}
	for _, s := range sections {
		if s.Slug == "" || s.Title == "" {
			t.Errorf("section with empty slug/title: %+v", s)
		}
		if seen["c:"+s.Slug] {
			t.Errorf("duplicate component slug %q", s.Slug)
		}
		seen["c:"+s.Slug] = true
		if len(s.Examples) == 0 {
			t.Errorf("%s: no examples", s.Slug)
		}
		for i, e := range s.Examples {
			if e.Component == nil {
				t.Errorf("%s: example %d has a nil component", s.Slug, i)
			}
		}
	}
	for _, p := range examplePages {
		if p.Slug == "" || p.Title == "" || p.Component == nil {
			t.Errorf("invalid example page: %+v", p)
		}
		if seen["e:"+p.Slug] {
			t.Errorf("duplicate example slug %q", p.Slug)
		}
		seen["e:"+p.Slug] = true
	}
}

// shadcnComponents is the full Components list from the shadcn/ui docs sidebar
// (https://ui.shadcn.com/docs/components), used to verify completeness of the
// port. Cross-referenced against github.com/shadcn-ui/ui.
var shadcnComponents = []string{
	"accordion", "alert", "alert-dialog", "aspect-ratio", "avatar", "badge",
	"breadcrumb", "button", "button-group", "calendar", "card", "carousel",
	"chart", "checkbox", "collapsible", "combobox", "command", "context-menu",
	"data-table", "date-picker", "dialog", "direction", "drawer", "dropdown-menu",
	"empty", "field", "hover-card", "input", "input-group", "input-otp", "item",
	"kbd", "label", "menubar", "native-select", "navigation-menu", "pagination",
	"popover", "progress", "radio-group", "resizable", "scroll-area", "select",
	"separator", "sheet", "sidebar", "skeleton", "slider", "sonner", "spinner",
	"switch", "table", "tabs", "textarea", "toast", "toggle", "toggle-group",
	"tooltip", "typography",
}

// TestMatchesShadcnDocs guarantees every component in the shadcn/ui docs is
// present in this port (a clean, complete port).
func TestMatchesShadcnDocs(t *testing.T) {
	have := map[string]bool{}
	for _, s := range sections {
		have[s.Slug] = true
	}
	var missing []string
	for _, slug := range shadcnComponents {
		if !have[slug] {
			missing = append(missing, slug)
		}
	}
	if len(missing) > 0 {
		t.Errorf("missing %d shadcn components: %v", len(missing), missing)
	}
	if len(sections) < len(shadcnComponents) {
		t.Errorf("registry has %d sections, expected at least %d", len(sections), len(shadcnComponents))
	}
}

// TestAllDemosHaveSource guarantees every component example can show its code
// (i.e. its Func resolves to embedded source).
func TestAllDemosHaveSource(t *testing.T) {
	for _, s := range sections {
		for i, ex := range s.Examples {
			if ex.Func == "" {
				t.Errorf("%s example %d has no Func", s.Slug, i)
				continue
			}
			if strings.TrimSpace(demoCode(ex.Func)) == "" {
				t.Errorf("%s: no embedded source for demo %q", s.Slug, ex.Func)
			}
		}
	}
}

// TestNoRenderArtifacts renders every component page and every example page and
// fails if the output contains tell-tale signs of a rendering bug: failed
// fmt.Sprintf verbs, stray "<nil>", templ's unsafe-content marker (ZgotmplZ,
// emitted when a bad URL/JS value is refused) or unrendered Go-template braces.
func TestNoRenderArtifacts(t *testing.T) {
	// Note: "{{" is intentionally not checked — code samples legitimately
	// contain Go slice-of-struct literals like []ChartDataPoint{{...}}.
	bad := []string{"%!", "<nil>", "ZgotmplZ"}
	render := func(t *testing.T, c templ.Component) string {
		var sb strings.Builder
		if err := c.Render(context.Background(), &sb); err != nil {
			t.Fatalf("render: %v", err)
		}
		return sb.String()
	}
	scan := func(t *testing.T, label, out string) {
		for _, b := range bad {
			if strings.Contains(out, b) {
				i := strings.Index(out, b)
				start := i - 60
				if start < 0 {
					start = 0
				}
				t.Errorf("%s: found artifact %q near: %q", label, b, out[start:i+len(b)])
			}
		}
	}

	cNav, eNav := componentNav(), examplesNav()
	for _, s := range sections {
		scan(t, "component/"+s.Slug, render(t, ComponentPage(cNav, s)))
	}
	for _, p := range examplePages {
		scan(t, "example/"+p.Slug, render(t, exampleDocument(eNav, p.Slug, p.Title, p.Component)))
	}
	scan(t, "home", render(t, Home(cNav)))
}

// TestBuildStaticSite exercises the static export and verifies the expected
// files are produced and links honor BasePath.
func TestBuildStaticSite(t *testing.T) {
	old := BasePath
	BasePath = "/shadcn-templ"
	defer func() { BasePath = old }()

	dir := t.TempDir()
	if err := buildStaticSite(dir); err != nil {
		t.Fatalf("buildStaticSite: %v", err)
	}

	want := []string{
		"index.html",
		".nojekyll",
		filepath.Join("static", "css", "styles.css"),
		filepath.Join("components", "button", "index.html"),
		filepath.Join("examples", "mail", "index.html"),
	}
	for _, rel := range want {
		info, err := os.Stat(filepath.Join(dir, rel))
		if err != nil {
			t.Errorf("missing %s: %v", rel, err)
			continue
		}
		if rel != ".nojekyll" && info.Size() == 0 {
			t.Errorf("%s is empty", rel)
		}
	}

	// Links must be prefixed with the base path for project-site hosting.
	index, err := os.ReadFile(filepath.Join(dir, "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(index), `href="/shadcn-templ/components/button"`) {
		t.Error("index.html links are not prefixed with BasePath")
	}
	if !strings.Contains(string(index), `href="/shadcn-templ/static/css/styles.css"`) {
		t.Error("index.html stylesheet link is not prefixed with BasePath")
	}
}
