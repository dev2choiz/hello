
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

docker-network: ## Create docker network
	docker network create hello || true

down: ## Stop docker containers
	docker-compose down --remove-orphans

up: ## Run docker container
	docker-compose up -d --build

start: docker-network down up ## start the stack

sh: ## Access to the hello-front-vue container
	docker-compose exec -u root hello-front-vue bash

## NPM
install: ## Run npm install in docker
	docker-compose exec hello-front-vue npm install

## VUE
serve: ## npm run serve
	docker-compose exec hello-front-vue npm run serve

gen-proto: ## Generate typescript protobuf
	docker-compose exec protoc ./scripts/generate_proto.sh
