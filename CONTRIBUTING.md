# Contributing to shadcn-templ

Thanks for your interest in contributing! This is a port of
[shadcn/ui](https://ui.shadcn.com/) to **Go + [Templ](https://templ.guide/) +
[Alpine.js](https://alpinejs.dev/) + Tailwind CSS** — server-rendered, no React
and no Node build step.

For the full API reference and authoring conventions, read
[`AGENTS.md`](AGENTS.md). For real usage of every component, see the demos in
[`cmd/examples/demos*.templ`](cmd/examples/). This guide is the short version.

## Prerequisites

- **Go 1.23+**
- **templ CLI** — `go install github.com/a-h/templ/cmd/templ@v0.3.1001`
  (match the version pinned in [`go.mod`](go.mod) / CI).
- **Tailwind standalone CLI** — download it once into the repo root (it's
  git-ignored); pick the binary for your platform from the
  [v3.4.17 release](https://github.com/tailwindlabs/tailwindcss/releases/tag/v3.4.17):

  ```sh
  curl -sL -o tailwindcss \
    https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-macos-arm64
  chmod +x tailwindcss
  ```

## Development loop

The [`Makefile`](Makefile) has everything you need:

| Command | What it does |
| --- | --- |
| `make gen` | `templ generate` + compile the CSS |
| `make tailwind` | Compile the CSS only |
| `make examples` | Run the docs/examples site at `:8080` |
| `make example` | Run the minimal starter app at `:8081` |
| `make site` | Export the static docs site to `./dist` |
| `make test` | Run the tests |
| `make verify` | `gen` + build + test + vet (run this before opening a PR) |

A typical session: edit a `.templ` file, run `make gen`, then `make examples`
to preview your change against the docs site.

## Generated files are committed

`.templ` files compile to `*_templ.go`, and Tailwind compiles to
`static/css/output.css`. **After editing any `.templ`, run `templ generate`
(or `make gen`) and `make tailwind`, then commit the regenerated `_templ.go`
and CSS alongside your source changes.** CI fails if the committed generated
code is stale (`git diff --exit-code` after `templ generate`).

## Authoring a component

- **One file per component** in `ui/`, **kebab-case** (e.g.
  `ui/radio-group.templ`). Components live in package `ui`; icons in `icons`.
- **Match shadcn/ui exactly** — the same Tailwind classes, roles, `aria-*`
  attributes and `data-state` values as the upstream component.
- **API contract:** every component's last two parameters are
  `classes string, attrs templ.Attributes`. Component-specific parameters come
  first. Merge the root `class` with `twmerge.Merge(baseClass, classes)` and
  spread `{ attrs... }` onto the root element so callers can override classes
  and add `id` / `hx-*` / `x-*` / ARIA attributes.
- Self-close void elements (`<input ... />`); escape a literal `@` in templ
  text with a Go string expr (`{ "@radix-ui" }`).
- Add a demo in `cmd/examples/demos*.templ` for any new component (the docs
  page renders it and shows its source). Tests cross-reference the registry
  against the shadcn/ui component list, so a missing entry fails CI.

## Before opening a pull request

1. Run **`make verify`** — it must pass (`templ generate` + build + test + vet).
2. Confirm the regenerated `_templ.go` and `output.css` are committed.
3. Fill out the pull request checklist.

Thanks again!
