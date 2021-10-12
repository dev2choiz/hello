
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

docker-network: ## Create docker network
	docker network create hello || true

down: ## Stop docker containers
	docker-compose down --remove-orphans

up: ## Run docker container
	docker-compose up -d --build

start: docker-network down up dev ## start the stack

sh: ## Access to the hello-front-next container
	docker-compose exec -u root hello-front-next bash
sh-protoc: ## Access to the protoc container
	docker-compose exec -u root protoc bash

gen-proto: ## Generate typescript protobuf
	docker-compose exec protoc ./scripts/generate_proto.sh

install: ## Run npm install in docker
	docker-compose exec hello-front-next npm install

logs-cdn: ## Display local cdn logs
	docker-compose logs -f hello-front-next-cdn

restart-protoc: ## Restart protoc container
	docker-compose stop protoc
	docker-compose up -d --build protoc
restart-dev-cdn: ## restart hello-front-next-dev-cdn container and display logs
	docker-compose stop hello-front-next-dev-cdn
	docker-compose up -d --build hello-front-next-dev-cdn
	docker-compose logs -f hello-front-next-dev-cdn
restart-prod: ## Restart the container running optimized app as in prod
	docker-compose stop hello-front-next-prod
	docker-compose up -d --build hello-front-next-prod
	docker-compose logs -f hello-front-next-prod

## next.js
next-clean: ## delete .next directory
	rm -rf ./.next
	rm -rf ./out
dev: ## Execute next with dev mode
	docker-compose exec hello-front-next npm run dev
next-build: next-clean
	docker-compose exec hello-front-next npm run build
next-start: next-build
	docker-compose exec hello-front-next npm run start
generate:
	docker-compose exec hello-front-next npm run generate
export: next-build
	docker-compose exec hello-front-next npm run export
