.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)



compose-down: ## Down docker compose
	docker compose down
.PHONY: compose-down

compose-up: compose-down ## Run docker compose
	docker compose up -d && docker compose logs --follow
.PHONY: compose-up

compose-build: compose-down ## Build docker compose
	docker compose build
.PHONY: compose-build

compose-build-up: compose-down compose-build compose-up ## Build and run docker compose
.PHONY: compose-build-up

swag-v1: ### swag init
	swag init -g internal/api/controller/http/v1/router.go -o docs
.PHONY: swag-v1