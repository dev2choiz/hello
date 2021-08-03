FROM golang:1.16-alpine as stage_base
WORKDIR /source

FROM stage_base as stage_dependencies
COPY . /source
RUN go mod vendor

FROM stage_dependencies as stage_builder
ARG APP_NAME
COPY . /source
RUN go mod vendor
WORKDIR "/source/cmd/apps/$APP_NAME"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o "/source/bin/$APP_NAME"
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o "/source/bin/$APP_NAME"
WORKDIR /source

FROM gcr.io/distroless/base as release
ARG APP_NAME
WORKDIR /
COPY --from=stage_builder "/source/bin/$APP_NAME" .

CMD ["/$APP_NAME"]
