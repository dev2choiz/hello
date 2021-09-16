FROM golang:1.16-alpine as stage_base
WORKDIR /source

FROM stage_base as stage_dependencies
COPY . /source
RUN go mod vendor

FROM stage_dependencies as stage_builder
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o "/source/bin/app"

FROM gcr.io/distroless/base as release
WORKDIR /
COPY --from=stage_builder "/source/bin/app" .
