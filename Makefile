
down:
	docker-compose down --remove-orphans

up:
	docker-compose up -d --build

start: down vendor up

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

deploy-endpoint: gen-proto
	gcloud endpoints services deploy api/config/api_descriptor.pb ./api/config/api_config.yaml

restart-api:
	docker-compose stop hello-api
	docker-compose up -d --build hello-api

restart-esp:
	docker-compose stop esp
	docker-compose up -d --build esp
