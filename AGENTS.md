# AGENTS.md — shadcn-templ

Guidance for AI assistants (Claude Code, Cursor, Copilot, …) working in or with
this repository. This is a port of [shadcn/ui](https://ui.shadcn.com) to **Go +
[Templ](https://templ.guide) + [Alpine.js](https://alpinejs.dev) + Tailwind
CSS**. There is no React and no JS build step: components render on the server
and Alpine.js adds interactivity in the browser.

Module path: `github.com/davidbudnick/shadcn-templ`

## Setup in a consuming app

1. Put `@shadcntempl.Head()` in your `<head>` (loads Alpine.js + the dark-mode /
   theme script).
2. Serve the embedded stylesheet `shadcntempl.CSS` and link it.

```go
import shadcntempl "github.com/davidbudnick/shadcn-templ"

mux.HandleFunc("GET /static/css/styles.css", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(shadcntempl.CSS)
})
```

```templ
templ Layout() {
	<!DOCTYPE html>
	<html lang="en">
		<head>
			@shadcntempl.Head()
			<link rel="stylesheet" href="/static/css/styles.css"/>
		</head>
		<body>{ children... }</body>
	</html>
}
```

## API conventions (read this before generating code)

- Components live in package `ui` (`github.com/davidbudnick/shadcn-templ/ui`).
  Icons live in package `icons`.
- **Every component's last two parameters are `classes string, attrs templ.Attributes`.**
  Component-specific parameters come first. Pass `""` and `nil` when unused.
- `classes` is merged with the component's base classes via tailwind-merge, so
  later utilities win (e.g. pass `"rounded-full"` to override `rounded-md`).
- `attrs` is spread onto the root element — use it for `id`, `hx-*`, `x-*`,
  `disabled`, ARIA, etc.
- Variants/sizes/types are typed constants, e.g. `ui.ButtonVariantOutline`,
  `ui.ButtonSizeIcon`, `ui.InputTypeEmail`.

```templ
@ui.Button(ui.ButtonTypeButton, ui.ButtonVariantOutline, ui.ButtonSizeDefault,
	"w-full", templ.Attributes{"hx-post": "/save"}) {
	Save
}
```

## Interactivity (Alpine.js)

Interactive components manage state with Alpine. Two patterns:

1. **Trigger-as-attributes**: some components expose an exported
   `templ.Attributes` you spread onto your own trigger element, e.g.
   `ui.DialogTriggerAttrs`, `ui.DropdownMenuTriggerAttrs`, `ui.PopoverTriggerAttrs`,
   `ui.SheetTriggerAttrs`, `ui.AlertDialogTriggerAttrs`, `ui.CollapsibleTriggerAttrs`.

```templ
@ui.Dialog("", nil) {
	@ui.Button(ui.ButtonTypeButton, ui.ButtonVariantOutline, ui.ButtonSizeDefault, "", ui.DialogTriggerAttrs) {
		Open
	}
	@ui.DialogContent("", nil) { ... }
}
```

2. **Explicit trigger components**: e.g. `ui.TooltipTrigger`, `ui.SelectTrigger`-style
   APIs, `ui.AccordionTrigger`. Follow the demos.

State variable names used internally are prefixed `shadcntempl_` (e.g.
`shadcntempl_open`) to avoid clashing with your own Alpine data.

## Theming

- **Dark mode**: toggle the `dark` class on `<html>` and persist
  `localStorage.theme`. `Head()` restores it before paint.
- **Color themes**: set `document.documentElement.dataset.theme` and persist
  `localStorage.themeColor` (`Head()` restores it). Base palettes:
  `slate stone gray neutral` (default/empty = Zinc); accents:
  `red rose orange green blue yellow violet`. Themes are CSS-variable based
  (`--primary`, `--ring`, …) — override the tokens in your own CSS to customize.

## Authoring new components (for contributors)

- One file per component in `ui/`, kebab-case (e.g. `ui/radio-group.templ`).
- Use `twmerge.Merge(baseClass, classes)` for the root `class` and spread
  `{ attrs... }`.
- Match shadcn/ui's exact Tailwind classes, roles, `aria-*` and `data-state`.
- Escape a literal `@` in templ **text** with a Go string expr: `{ "@radix-ui" }`.
- Self-close void elements: `<input ... />`.
- After editing `.templ`, run `templ generate` and `make tailwind`.

## Where to look

- `ui/*.templ` — the components.
- `cmd/examples/demos*.templ` — a demo per component (the source shown on each
  docs page). The clearest reference for how to use each component.
- `example/` — a minimal standalone app consuming the library.
- `make examples` runs the docs site; `make example` runs the starter app.

## Verify your changes

```sh
make verify   # templ generate + go build + go test + go vet
```

Tests render every component and example, and `TestMatchesShadcnDocs` checks the
registry covers the full shadcn/ui component list.
