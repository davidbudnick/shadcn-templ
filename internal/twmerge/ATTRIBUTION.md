# Vendored: tailwind-merge-go

`internal/twmerge`, `internal/cache` and `internal/lru` are a vendored copy of
[github.com/Oudwins/tailwind-merge-go](https://github.com/Oudwins/tailwind-merge-go)
(v0.2.1, MIT License — see `LICENSE`).

It is copied into the module on purpose: the project deliberately depends only on
`github.com/a-h/templ`, so this utility lives in-tree rather than as an external
module (no third-party supply chain). Only the import paths were changed to point
at the internal packages.
