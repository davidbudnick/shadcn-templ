# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.3.0] - 2026-05-26

### Changed

- **All 59 components updated to match current shadcn/ui (Tailwind v4 /
  new-york).** Every element now carries a `data-slot` attribute, and the
  class set was modernized: the v4 focus ring (`focus-visible:ring-[3px]`
  `ring-ring/50` + `aria-invalid:*` error styling), `size-*` utilities,
  `shadow-xs`, and `transition-[color,box-shadow]`. The embedded stylesheet
  (`shadcntempl.CSS`) is regenerated accordingly. Component function APIs are
  unchanged; rendered markup/styling now tracks upstream more closely.

### Added

- **Accessibility & keyboard support** across interactive components:
  `aria-labelledby`/`aria-describedby` wiring and focus management for
  dialog/alert-dialog/sheet/drawer/popover; roving tabindex + arrow keys for
  tabs and radio-group; Space-to-toggle for checkbox; type-ahead-free arrow
  navigation, submenus and checkbox/radio items for dropdown/context/menubar;
  keyboard nav for command; `role="status"`/`aria-live` for toast and sonner;
  end-disabling + arrow keys for carousel; `role="grid"` semantics for calendar.
- New sub-components to match shadcn/ui: Field (Group/Set/Legend/Content/
  Title/Separator), Item (variants/sizes + Group/Header/Footer/Separator),
  Empty (Header/Media/Content), Kbd (Group), and Sub/Checkbox/Radio items for
  the menu components.
- Docs site: a global **⌘K / Ctrl+K command palette** searching every
  component and example, a **"Components" dropdown menu** in the nav, fully
  mobile-friendly example pages, and a clean SVG favicon.

### Fixed

- Renamed the misspelled `AccorionContent` component to `AccordionContent`
  (a deprecated `AccorionContent` alias is kept so existing callers compile).
- `Separator` now spreads `attrs` onto its root; fixed an invalid
  `botton-full` class in the dropdown menu.
- Calendar now gives each day a correct `aria-label` and navigates months
  client-side (no full page reload); data-table row selection is keyed by a
  stable id and exposes `aria-sort`.
- The docs component preview no longer clips overlays (calendar, popovers,
  dropdowns), and the desktop "Components" dropdown renders at full width.
- CI runs the race detector only on the latest Go version, balancing the
  build matrix.

## [0.2.0] - 2026-05-25

### Changed

- **Migrated styling to Tailwind CSS v4** (v4.3.0). Tailwind is now configured
  CSS-first in `static/css/input.css` via `@theme` / `@plugin` / `@source`;
  there is no `tailwind.config.js`. The embedded stylesheet (`shadcntempl.CSS`)
  is regenerated, but component markup and emitted classes are unchanged.
- **Raised the minimum Go version to 1.25** (previously documented as 1.23),
  required by the templ upgrade below.
- Bumped `github.com/a-h/templ` to `v0.3.1020` and regenerated every component
  (its codegen now uses the `templ.ResolveAttributeValue` runtime API).
- Bumped pinned GitHub Actions: `checkout` v6, `upload-artifact` v7,
  `configure-pages` v6, `upload-pages-artifact` v5, `deploy-pages` v5.

### Added

- CI now builds, vets, and race-tests across Go 1.25 and 1.26 and uploads a
  per-version coverage profile as a workflow artifact.
- Repository metadata: Dependabot config, `CONTRIBUTING.md`, `SECURITY.md`,
  issue and pull-request templates, `.editorconfig`, and this changelog.

### Removed

- `tailwind.config.js` — replaced by the CSS-first `@theme` configuration in
  `static/css/input.css`.

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

[0.3.0]: https://github.com/davidbudnick/shadcn-templ/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/davidbudnick/shadcn-templ/compare/v0.1.1...v0.2.0
[0.1.1]: https://github.com/davidbudnick/shadcn-templ/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/davidbudnick/shadcn-templ/releases/tag/v0.1.0
