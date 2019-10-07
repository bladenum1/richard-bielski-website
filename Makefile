all: build_go build_js build_html serve

build_go:
	tinygo build -o dist/dev/wasm/wasm.wasm -target wasm ./cmd/go_sdk

build_js:
	npx webpack --mode development

build_html:
	cp cmd/html/*.html dist/dev

serve:
	cd dist/dev && serve

init:
	mkdir dist
	mkdir dist/dev
	mkdir dist/dev/js
	mkdir dist/dev/wasm