package shadcntempl

import (
	"context"
	"strings"
	"testing"
)

func TestCSSEmbedded(t *testing.T) {
	if len(CSS) == 0 {
		t.Fatal("embedded CSS is empty; run `make tailwind` to compile static/css/output.css")
	}
}

func TestHeadRenders(t *testing.T) {
	var sb strings.Builder
	if err := Head().Render(context.Background(), &sb); err != nil {
		t.Fatalf("render Head: %v", err)
	}
	out := sb.String()
	for _, want := range []string{"alpinejs", "localStorage.theme", "themeColor", "dataset.theme"} {
		if !strings.Contains(out, want) {
			t.Errorf("Head output missing %q", want)
		}
	}
}

// TestCSSHasThemes checks the embedded stylesheet ships the shadcn color theme
// presets and the core semantic tokens.
func TestCSSHasThemes(t *testing.T) {
	css := string(CSS)
	for _, want := range []string{"--primary", "--ring", "--background"} {
		if !strings.Contains(css, want) {
			t.Errorf("CSS missing token %q", want)
		}
	}
	// Minification may drop the quotes in attribute selectors, so match the
	// unquoted form that survives in the compiled output.
	themes := []string{
		// accent colors
		"rose", "blue", "green", "violet", "red", "orange", "yellow",
		// base color palettes
		"slate", "stone", "gray", "neutral",
	}
	for _, theme := range themes {
		if !strings.Contains(css, "data-theme="+theme) {
			t.Errorf("CSS missing color theme %q", theme)
		}
	}
}
