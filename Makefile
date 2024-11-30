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

.PHONY: test
test:
	go test -v ./...

.PHONY: mockgen
mockgen:
	mockgen -source=../internal/domain/issue_commenter.go -destination=../internal/domain/mock/issue_commenter.go -package=domainmock

.PHONY: install-tools
install-tools:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: oapigen
oapigen:
	oapi-codegen -config ./openapi/config.yaml ./openapi/schema.yaml