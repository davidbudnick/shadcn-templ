package ui

import (
	"strings"
)

// selectDataExpr builds the Alpine x-data object for a Select, including a
// value->label map so the trigger can display the selected option's label
// without relying on querying rendered DOM text (which breaks for rich options).
func selectDataExpr(alpineVar string, options []SelectOption) string {
	var b strings.Builder
	b.WriteString("{ ")
	b.WriteString(alpineVar)
	b.WriteString(": '', ")
	b.WriteString(alpineVar)
	b.WriteString("_open: false, ")
	b.WriteString(alpineVar)
	b.WriteString("_labels: {")
	for i, opt := range options {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(jsStringLiteral(opt.Value))
		b.WriteString(": ")
		b.WriteString(jsStringLiteral(opt.Label))
	}
	b.WriteString("} }")
	return b.String()
}

// jsStringLiteral returns a single-quoted JS string literal with the necessary
// characters escaped.
func jsStringLiteral(s string) string {
	r := strings.NewReplacer(
		`\`, `\\`,
		`'`, `\'`,
		"\n", `\n`,
		"\r", `\r`,
	)
	return "'" + r.Replace(s) + "'"
}
