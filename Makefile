
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## Docker
docker-network: ## Create a docker network
	docker network create hello || true

down: ## Stop docker containers
	docker-compose down --remove-orphans

up: ## Start docker containers
	docker-compose up -d --build

start: docker-network down vendor up migration-init migration-up  ## Start docker containers, install go dependencies and run micro-services

debug: ## Run with debug mode using delve
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml stop hello-api
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml up -d --build hello-api
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml logs -f hello-api

logs: ## Show the main api logs
	docker-compose -f docker-compose.yml logs -f hello-api
logs1: ## Show the svc1 logs
	docker-compose -f docker-compose.yml logs -f hello-svc1
logs2: ## Show the svc2 logs
	docker-compose -f docker-compose.yml logs -f hello-svc2
logs-esp: ## Show the esp logs
	docker-compose -f docker-compose.yml logs -f esp

test:
	go test -v ./...

sh: ## Access to the api container
	docker-compose exec hello-api bash
sh-proto: ## Access to the proto container
	docker-compose exec protoc bash

vendor: ## Generate go dependencies
	docker run -v `pwd`:`pwd` -w `pwd` golang:1.17 go mod tidy
	docker run -v `pwd`:`pwd` -w `pwd` golang:1.17 go mod vendor

tag: ## example: TAG=v1.0.1 make tag
	scripts/git-tag.sh $(TAG)

gen-proto: ## Generate go protobuf files
	docker-compose exec -w /app protoc ./scripts/generate_proto.sh

## Restart container
restart-api: ## Restart the api container
	docker-compose stop hello-api
	docker-compose up -d --build hello-api
restart-svc1:  ## Restart the svc1 container
	docker-compose stop hello-svc1
	docker-compose up -d --build hello-svc1
restart-svc2:  ## Restart the svc2 container
	docker-compose stop hello-svc2
	docker-compose up -d --build hello-svc2
restart-esp:  ## Restart the esp container
	docker-compose stop esp
	docker-compose up -d --build esp
restart-protoc:  ## Restart the protoc container
	docker-compose stop protoc
	docker-compose up -d --build protoc
restart-proxy:  ## Restart the sql-proxy container
	docker-compose stop sql-proxy
	docker-compose up -d --build sql-proxy
restart-postgres:  ## Restart the postgres container
	docker-compose stop postgres
	docker-compose up -d --build postgres

## DEV COMMANDS
vendor-dev:  ## Restart the api container
	docker-compose exec -w /app/cmd/dev hello-api go mod vendor
#### Migrations
migration-diff: vendor-dev ## Generate a diff migration
	docker-compose exec -u hello -w /app/cmd/dev hello-api go run *.go migration diff
migration-init: vendor-dev  ## Init the table migration
	docker-compose exec -u hello -w /app/cmd/dev hello-api go run . migration init
migration-up: vendor-dev ## Apply migrations
	docker-compose exec -u hello -w /app/cmd/dev hello-api go run . migration up
	#docker-compose exec -u hello -w /app/cmd/dev hello-api dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient debug . -- migration up


## Infra
clean-pods: ## Delete shutting down pods (because of preemptive vm)
	kubectl get pods --all-namespaces | grep -i shutdown && kubectl get pods --all-namespaces | grep -i shutdown | awk '{print $$1, $$2}' | xargs kubectl delete pod -n

deploy-endpoint: gen-proto ## deploy esp in gcp
	gcloud endpoints services deploy api/config/api_descriptor.pb ./api/config/api_config.yaml
