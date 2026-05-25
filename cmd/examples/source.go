package main

import (
	"embed"
	"strings"
)

// demoFiles embeds the demo .templ sources so each component page can display
// the exact Go/templ code that produces its preview.
//
//go:embed demos.templ demos_grpa.templ demos_grpb.templ demos_grpc.templ demos_grpd.templ demos_grpe.templ demos_grpf.templ demos_extra.templ
var demoFiles embed.FS

// demoSources maps a templ function name (e.g. "ButtonDemo") to its source.
var demoSources = loadDemoSources()

func loadDemoSources() map[string]string {
	out := map[string]string{}
	entries, _ := demoFiles.ReadDir(".")
	for _, e := range entries {
		b, err := demoFiles.ReadFile(e.Name())
		if err != nil {
			continue
		}
		extractFuncs(string(b), out)
	}
	return out
}

// extractFuncs pulls each top-level `templ Name(...) { ... }` block out of a
// source file. templ functions close with a `}` in the first column.
func extractFuncs(src string, out map[string]string) {
	lines := strings.Split(src, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if !strings.HasPrefix(line, "templ ") {
			continue
		}
		name := strings.TrimSpace(line[len("templ "):])
		if idx := strings.IndexByte(name, '('); idx >= 0 {
			name = name[:idx]
		}
		block := []string{line}
		j := i + 1
		for ; j < len(lines); j++ {
			block = append(block, lines[j])
			if lines[j] == "}" {
				break
			}
		}
		out[name] = strings.Join(block, "\n")
		i = j
	}
}

// demoCode returns the source of the named demo function, or "" if unknown.
func demoCode(fn string) string {
	return demoSources[fn]
}
