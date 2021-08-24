
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## Docker
docker-network:
	docker network create hello || true

down:
	docker-compose down --remove-orphans

up:
	docker-compose up -d --build

start: docker-network down vendor up

debug:
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml stop hello-api
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml up -d --build hello-api
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml logs -f hello-api

logs:
	docker-compose -f docker-compose.yml logs -f hello-api
logs1:
	docker-compose -f docker-compose.yml logs -f hello-svc1
logs2:
	docker-compose -f docker-compose.yml logs -f hello-svc2
logs-esp:
	docker-compose -f docker-compose.yml logs -f esp

test:
	go test -v ./...

sh:
	docker-compose exec hello-api bash
sh-proto:
	docker-compose exec protoc bash

vendor:
	docker run -v `pwd`:`pwd` -w `pwd` golang:1.16 go mod tidy
	docker run -v `pwd`:`pwd` -w `pwd` golang:1.16 go mod vendor

tag: ## example: TAG=v1.0.1 make tag
	scripts/git-tag.sh $(TAG)

gen-proto:
	docker-compose exec -w /app protoc ./scripts/generate_proto.sh

## Restart container
restart-api:
	docker-compose stop hello-api
	docker-compose up -d --build hello-api
restart-esp:
	docker-compose stop esp
	docker-compose up -d --build esp
restart-protoc:
	docker-compose stop protoc
	docker-compose up -d --build protoc

## Infra
clean-pods:
	kubectl get pods --all-namespaces | grep -i shutdown && kubectl get pods --all-namespaces | grep -i shutdown | awk '{print $$1, $$2}' | xargs kubectl delete pod -n

deploy-endpoint: gen-proto
	gcloud endpoints services deploy api/config/api_descriptor.pb ./api/config/api_config.yaml
