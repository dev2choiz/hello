
down:
	docker-compose down --remove-orphans

up:
	docker-compose up -d --build

start: down up

sh:
	docker-compose exec -u root hello-front bash
sh-protoc:
	docker-compose exec -u root protoc bash

gen-proto:
	docker-compose exec protoc ls -la
	docker-compose exec protoc ./scripts/generate_proto.sh

install:
	docker-compose exec hello-front npm install

dev:
	docker-compose exec hello-front npm run dev

restart-protoc:
	docker-compose stop protoc
	docker-compose up -d --build protoc
