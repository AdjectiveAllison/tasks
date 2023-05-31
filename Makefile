.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@latest
	tinygo build -o ./build/app.wasm -target wasm ./main.go
# TODO: Consider changing the above line to below -- but idk if it will break.
# tinygo build -o ./build/app.wasm -target wasm ./...

.PHONY: deploy
deploy:
	wrangler publish
#maybe this needs to be 'wrangler deploy' instead of publish if I was on the correct version?

.PHONY: generate
generate:
	go generate ./...

.PHONY: init-db
init-db:
	wrangler d1 execute tasks --file=./schema.sql
	wrangler d1 execute tasks --file=./tasks.sql
# TODO: Remove this execution of tasks.sql

.PHONY: init-db-preview
init-db-preview:
	wrangler d1 execute tasks-preview --file=./schema.sql