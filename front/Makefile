
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

docker-network:
	docker network create hello || true

down:
	docker-compose down --remove-orphans

up:
	docker-compose up -d --build

start: docker-network down up dev

sh:
	docker-compose exec -u root hello-front bash
sh-protoc:
	docker-compose exec -u root protoc bash

gen-proto:
	docker-compose exec protoc ./scripts/generate_proto.sh

install:
	docker-compose exec hello-front npm install

logs-cdn:
	docker-compose logs -f hello-front-cdn

restart-protoc:
	docker-compose stop protoc
	docker-compose up -d --build protoc
restart-dev-cdn:
	docker-compose stop hello-front-dev-cdn
	docker-compose up -d --build hello-front-dev-cdn
	docker-compose logs -f hello-front-dev-cdn
restart-prod:
	docker-compose stop hello-front-prod
	docker-compose up -d --build hello-front-prod
	docker-compose logs -f hello-front-prod

## next.js
next-clean: ## delete .next directory
	rm -rf ./.next
	rm -rf ./out
dev:
	docker-compose exec hello-front npm run dev
next-build: next-clean
	docker-compose exec hello-front npm run build
next-start: next-build
	docker-compose exec hello-front npm run start
generate:
	docker-compose exec hello-front npm run generate
export: next-build
	docker-compose exec hello-front npm run export
