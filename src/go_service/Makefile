gvw:
	@templ generate
	@pnpm run style

build:
	@templ generate
	@pnpm run style
	@go build -o bin/gore.exe main.go

run: build
	@./bin/gore.exe

air:
	@air
