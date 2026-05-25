# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2026-05-24

### Added

- `icons.Menu` (Lucide hamburger) icon.
- Responsive mobile navigation for the docs site: a hamburger menu in the
  header that toggles an Alpine.js-powered drawer listing the Examples and
  Components links.

### Changed

- Made the docs site layout responsive on small screens — the top nav collapses
  on mobile, component previews and code blocks scroll horizontally instead of
  overflowing, and the main content area no longer pushes the layout wide.

## [0.1.0] - 2026-05-24

### Added

- Initial release: a complete port of every shadcn/ui component (59) plus all
  shadcn/ui page examples (Authentication, Dashboard, Cards, Tasks, Playground,
  Forms, Music, Mail), rendered server-side with Go + Templ + Alpine.js +
  Tailwind CSS.
- Consistent component API — every component takes a trailing
  `classes string` (conflict-merged with the base classes via the vendored
  tailwind-merge) and `attrs templ.Attributes` spread onto the root element.
- Dark mode and shadcn/ui accent color themes (Zinc, Red, Rose, Orange, Green,
  Blue, Yellow, Violet), CSS-variable based and persisted in `localStorage`.
- Embedded compiled stylesheet (`shadcntempl.CSS`) and a `Head()` component
  that loads Alpine.js and the theme script.
- A documentation site with a live preview and source for every component,
  plus a static-site generator for GitHub Pages.
- CI: build / vet / race tests with a generated-code freshness check, and a
  GitHub Pages deploy workflow.
- `AGENTS.md` skill file documenting the API and conventions for AI assistants.

[Unreleased]: https://github.com/davidbudnick/shadcn-templ/compare/v0.1.1...HEAD
[0.1.1]: https://github.com/davidbudnick/shadcn-templ/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/davidbudnick/shadcn-templ/releases/tag/v0.1.0
