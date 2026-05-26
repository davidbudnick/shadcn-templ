package main

import (
	"strings"

	"github.com/a-h/templ"
)

// faviconDataURL returns a clean, self-contained SVG favicon as a data URL: a
// rounded dark square with a white lowercase "s" wordmark. Inlining it avoids an
// extra asset/route and works for both the live server and the static export.
func faviconDataURL() templ.SafeURL {
	const svg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32">` +
		`<rect width="32" height="32" rx="7" fill="#18181b"/>` +
		`<text x="16" y="23" text-anchor="middle" font-family="ui-sans-serif,system-ui,-apple-system,Segoe UI,Roboto,sans-serif" font-size="21" font-weight="700" fill="#fafafa">s</text>` +
		`</svg>`
	enc := strings.NewReplacer(
		`"`, "%22",
		"#", "%23",
		"<", "%3C",
		">", "%3E",
		" ", "%20",
		",", "%2C",
	).Replace(svg)
	return templ.SafeURL("data:image/svg+xml," + enc)
}
