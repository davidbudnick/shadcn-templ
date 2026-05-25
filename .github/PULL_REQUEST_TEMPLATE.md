<!--
Thanks for contributing! Please fill out the checklist below.
See CONTRIBUTING.md for the full workflow.
-->

## Summary

<!-- What does this PR do, and why? Link any related issue (e.g. Closes #123). -->

## Checklist

- [ ] `make verify` passes locally (`templ generate` + build + test + vet).
- [ ] Regenerated `templ` Go code and the Tailwind CSS, and committed the
      generated files (`*_templ.go`, `static/css/output.css`).
- [ ] Markup matches shadcn/ui — same Tailwind classes, `aria-*` and
      `data-state` as the upstream component.
- [ ] Followed the API contract: last two parameters are
      `classes string, attrs templ.Attributes`, merged/spread on the root.
- [ ] If adding or changing a component: updated the demo in
      `cmd/examples/demos*.templ` (and docs as needed).
