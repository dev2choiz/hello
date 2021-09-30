FROM golang:1.17 as stage_base
WORKDIR /app

FROM stage_base as stage_dev
RUN useradd -ms /bin/bash hello

FROM stage_dev as stage_debug
RUN go install github.com/go-delve/delve/cmd/dlv@latest
