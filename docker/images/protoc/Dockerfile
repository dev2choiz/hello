FROM golang:1.16 as go_dependencies
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

FROM debian:buster-slim as protoc_builder
ARG VERS=3.17.3
ARG ARCH=linux-x86_64
RUN echo "Building protoc Cloud Builder ${VERS}-${ARCH}" && \
    apt-get update -y && apt-get upgrade -y && \
    apt-get install wget unzip -y && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/* && \
    wget "https://github.com/protocolbuffers/protobuf/releases/download/v${VERS}/protoc-${VERS}-${ARCH}.zip" && \
    unzip "protoc-${VERS}-${ARCH}.zip" -d protoc && \
    rm "protoc-${VERS}-${ARCH}.zip"

FROM debian:buster-slim  as release
COPY --from=protoc_builder /protoc /protoc
COPY --from=go_dependencies "/go/bin/protoc-gen-go" /utils/bin/
COPY --from=go_dependencies "/go/bin/protoc-gen-go-grpc" /utils/bin/
COPY --from=go_dependencies "/go/bin/protoc-gen-grpc-gateway" /utils/bin/
COPY --from=go_dependencies "/go/bin/protoc-gen-openapiv2" /utils/bin/
ENV PATH=$PATH:/protoc/bin/:/utils/bin/

FROM release  as release_dev
RUN useradd -ms /bin/bash nonroot
RUN chown -R nonroot:nonroot /protoc/bin/ /utils/bin/
USER nonroot
