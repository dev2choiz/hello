FROM node:16-alpine as next_builder
WORKDIR /app/front
USER node
COPY --chown=node:node . /app/front
RUN npm install
RUN npm build

FROM next_builder as release
WORKDIR /app/front
COPY --from=next_builder /app/front /app/front
