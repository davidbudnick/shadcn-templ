package main

import "encoding/json"

// searchEntry is one row in the client-side command-palette index.
type searchEntry struct {
	T string `json:"t"` // title
	U string `json:"u"` // url
	G string `json:"g"` // group label ("Component" / "Example")
}

// searchIndexJSON builds the JSON the Cmd/Ctrl+K palette filters over: every
// component and example page, with its title, URL and group. json.Marshal
// escapes <, > and & so the result is safe to embed directly in a <script>.
func searchIndexJSON() string {
	var entries []searchEntry
	for _, n := range componentNav() {
		entries = append(entries, searchEntry{T: n.Title, U: string(urlComponent(n.Slug)), G: "Component"})
	}
	for _, n := range examplesNav() {
		entries = append(entries, searchEntry{T: n.Title, U: string(urlExample(n.Slug)), G: "Example"})
	}
	b, err := json.Marshal(entries)
	if err != nil {
		return "[]"
	}
	return string(b)
}
