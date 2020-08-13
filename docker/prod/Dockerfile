FROM golang:latest as hello_api_base
WORKDIR /source

FROM hello_api_base as hello_api_builder
COPY . /source
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./bin/app
RUN pwd && ls bin -la

FROM gcr.io/distroless/base as hello_api_release
WORKDIR /app
COPY --from=hello_api_builder /source/bin/app /app/

