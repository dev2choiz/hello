FROM golang:1.17 as stage_base
WORKDIR /app

FROM stage_base as stage_dev
#RUN go get github.com/t-yuki/gocover-cobertura
#RUN go get github.com/stretchr/testify

FROM stage_dev as stage_debug
RUN go install github.com/go-delve/delve/cmd/dlv@latest
