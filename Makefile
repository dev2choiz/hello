
down:
	docker-compose down --remove-orphans

up:
	docker-compose up -d --build

test:
	go test -v ./...

sh:
	docker-compose exec hello bash

build-app:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/app
