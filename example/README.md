# Example app

A self-contained app that consumes [`shadcn-templ`](../) — a faithful copy of
the [shadcn/ui Dashboard example](https://ui.shadcn.com/examples/dashboard)
built entirely from library components (top nav with search + Avatar + Dropdown
Menu, Tabs, stat Cards, an overview chart and a recent-sales list) with dark
mode. A realistic starting point for your own project.

## Run

From the repo root:

```sh
go run ./example
```

Then open <http://localhost:8081>.

## How it works

- [`main.go`](main.go) serves two routes: the page at `/` and the library's
  embedded stylesheet at `/static/css/styles.css` (`shadcntempl.CSS`).
- [`page.templ`](page.templ) puts `@shadcntempl.Head()` in `<head>` (which loads
  Alpine.js and the dark-mode script) and composes the page from `ui.*`
  components.

## A note on CSS

This example reuses the stylesheet embedded in the library, which already
contains every class the components use. In your own project you would run the
[Tailwind CLI](https://tailwindcss.com/) over your templates **and** the
library source so that any extra utility classes you add are included:

```js
// tailwind.config.js
content: [
  "./**/*.templ",
  "./vendor/github.com/davidbudnick/shadcn-templ/**/*.templ",
],
```
