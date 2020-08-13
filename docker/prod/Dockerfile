FROM golang:1.16-alpine as hello_api_base
WORKDIR /source

FROM hello_api_base as hello_api_builder
COPY . /source
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./bin/app
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./bin/app

FROM gcr.io/distroless/base as hello_api_release
WORKDIR /
COPY --from=hello_api_builder /source/bin/app .

CMD ["/app"]
