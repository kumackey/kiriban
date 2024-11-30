.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.23.1
	tinygo build -o ./build/app.wasm -target wasm -no-debug ./cmd/api/main.go

.PHONY: deploy
deploy:
	wrangler deploy
