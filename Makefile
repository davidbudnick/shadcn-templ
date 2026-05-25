.PHONY: gen examples example site tailwind tidy test build verify

# Generate templ Go code and compile the Tailwind CSS.
gen:
	templ generate
	./tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

# Compile only the CSS.
tailwind:
	./tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

# Run the docs/examples site (component demos + full page examples).
examples:
	go run ./cmd/examples

# Run the minimal example consumer app (a feedback form).
example:
	go run ./example

# Export the docs as a static site into ./dist (hostable on GitHub Pages).
# Pass BASE to set a sub-path, e.g. `make site BASE=/shadcn-templ`.
BASE ?=
site: gen
	go run ./cmd/examples -build ./dist -base $(BASE)

tidy:
	go mod tidy

test:
	go test ./...

build:
	go build ./...

# Full pipeline used in CI and locally before committing.
verify: gen build test
	go vet ./...
