package main

import "strings"

// A tiny, dependency-free syntax highlighter for the templ/Go code samples.
// templ is a superset of Go, so a small tokenizer (comments, strings, numbers,
// keywords and call names) is enough to look good — and it keeps the module
// free of any external lexer dependency. Colors are emitted as inline styles
// (github-dark palette) so no extra stylesheet is needed and it works in the
// static export.

var hlKeywords = map[string]bool{
	"templ": true, "package": true, "import": true, "func": true, "if": true,
	"else": true, "for": true, "range": true, "return": true, "var": true,
	"const": true, "type": true, "struct": true, "interface": true, "map": true,
	"switch": true, "case": true, "default": true, "go": true, "true": true,
	"false": true, "nil": true, "break": true, "continue": true, "chan": true,
}

const (
	colKeyword = "#ff7b72"
	colString  = "#a5d6ff"
	colComment = "#8b949e"
	colNumber  = "#79c0ff"
	colFunc    = "#d2a8ff"
	colText    = "#e6edf3"
)

func htmlEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

func hlSpan(color, text string) string {
	return `<span style="color:` + color + `">` + htmlEscape(text) + `</span>`
}

func isIdentStart(c rune) bool {
	return c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isIdentPart(c rune) bool {
	return isIdentStart(c) || (c >= '0' && c <= '9')
}

// highlight tokenizes templ/Go source into syntax-highlighted HTML.
func highlight(code string) string {
	var b strings.Builder
	b.WriteString(`<pre style="color:` + colText + `;margin:0"><code>`)
	r := []rune(code)
	n := len(r)
	for i := 0; i < n; {
		c := r[i]
		switch {
		case c == '/' && i+1 < n && r[i+1] == '/':
			j := i
			for j < n && r[j] != '\n' {
				j++
			}
			b.WriteString(hlSpan(colComment, string(r[i:j])))
			i = j
		case c == '"' || c == '`' || c == '\'':
			quote := c
			j := i + 1
			for j < n {
				if r[j] == '\\' && quote != '`' {
					j += 2
					continue
				}
				if r[j] == quote {
					j++
					break
				}
				j++
			}
			if j > n {
				j = n
			}
			b.WriteString(hlSpan(colString, string(r[i:j])))
			i = j
		case isIdentStart(c):
			j := i
			for j < n && isIdentPart(r[j]) {
				j++
			}
			word := string(r[i:j])
			switch {
			case hlKeywords[word]:
				b.WriteString(hlSpan(colKeyword, word))
			case j < n && r[j] == '(':
				b.WriteString(hlSpan(colFunc, word))
			default:
				b.WriteString(htmlEscape(word))
			}
			i = j
		case c >= '0' && c <= '9':
			j := i
			for j < n && ((r[j] >= '0' && r[j] <= '9') || r[j] == '.' || r[j] == 'x' || (r[j] >= 'a' && r[j] <= 'f') || (r[j] >= 'A' && r[j] <= 'F')) {
				j++
			}
			b.WriteString(hlSpan(colNumber, string(r[i:j])))
			i = j
		default:
			b.WriteString(htmlEscape(string(c)))
			i++
		}
	}
	b.WriteString(`</code></pre>`)
	return b.String()
}

// demoHTML is the highlighted HTML for each demo, computed once.
var demoHTML = func() map[string]string {
	out := make(map[string]string, len(demoSources))
	for fn, src := range demoSources {
		out[fn] = highlight(src)
	}
	return out
}()

// demoCodeHTML returns highlighted HTML for the named demo, or "" if unknown.
func demoCodeHTML(fn string) string {
	return demoHTML[fn]
}
