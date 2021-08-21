FROM golang:1.16-alpine as stage_base
WORKDIR /source

FROM stage_base as stage_dependencies
COPY . /source
RUN go mod vendor

FROM stage_dependencies as stage_builder
## Build hello-api
WORKDIR "/source/cmd/apps/hello-api"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o "/source/bin/hello-api"
## Build hello-svc
WORKDIR "/source/cmd/apps/hello-svc"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o "/source/bin/hello-svc"
WORKDIR /source

FROM gcr.io/distroless/base as release
WORKDIR /
COPY --from=stage_builder "/source/bin/hello-api" .
COPY --from=stage_builder "/source/bin/hello-svc" .
