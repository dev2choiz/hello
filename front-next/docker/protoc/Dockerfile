FROM debian:buster-slim as protoc_builder
ARG VERS=3.17.3
ARG ARCH=linux-x86_64
RUN apt-get update && apt-get install wget unzip -y \
    && rm -rf /var/lib/apt/lists/* \
    && wget "https://github.com/protocolbuffers/protobuf/releases/download/v${VERS}/protoc-${VERS}-${ARCH}.zip" \
    && unzip "protoc-${VERS}-${ARCH}.zip" -d protoc \
    && rm "protoc-${VERS}-${ARCH}.zip"

FROM node:16-buster-slim as release
RUN npm i -g grpc-tools ts-protoc-gen
COPY --from=protoc_builder /protoc /protoc
ENV PATH=$PATH:/protoc/bin
WORKDIR /app/front-next

FROM release as release_dev
RUN chown -R node:node /protoc/
USER node
