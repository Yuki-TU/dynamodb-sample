# makeを打った時のコマンド
.DEFAULT_GOAL := help

codegen: openapi ## コード自動生成
	oapi-codegen -package gen -generate types -o gen/types.gen.go ./docs/openapi.yml
	oapi-codegen -package gen -generate strict-server,gin -templates ./_tools/oapi/templates -o gen/server.gen.go ./docs/openapi.yml

.PHONY: openapi
openapi: ## OpenAPI bundle
	redocly bundle ./docs/index.openapi.yml --output ./docs/openapi.yml

.PHONY: log
log: ## ログ確認
	docker compose logs -f app

.PHONY: logdb
logdb: ## DBログ確認
	docker compose logs -f dynamodb

in: ## コンテナに入る
	docker compose exec app sh

.PHONY: help
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
