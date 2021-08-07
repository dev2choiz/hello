
down:
	docker-compose down --remove-orphans

up:
	docker-compose up -d --build

start: down vendor up

debug:
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml stop primary
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml up -d --build primary
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml logs -f primary

logs:
	docker-compose -f docker-compose.yml logs -f primary

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
