
down:
	docker-compose down --remove-orphans

up:
	docker-compose up -d --build

debug:
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml stop hello
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml up -d --build hello
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml logs -f hello

logs:
	docker-compose -f docker-compose.yml logs -f hello

test:
	go test -v ./...

sh:
	docker-compose exec hello bash

build-app:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/app
