gvw:
	@templ generate
	@pnpm run style

build:
	@templ generate
	@pnpm run style
	@go build -o bin/bff.exe main.go

run: build
	@./bin/bff.exe

air:
	@air
