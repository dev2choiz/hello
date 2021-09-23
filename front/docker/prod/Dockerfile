FROM node:16-alpine as next_builder
WORKDIR /app/front
COPY --chown=node:node . /app/front
RUN npm install
RUN npm build

FROM next_builder as release
USER node
WORKDIR /app/front
COPY --from=next_builder --chown=node:node /app/front /app/front
