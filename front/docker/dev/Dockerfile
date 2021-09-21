FROM node:16 as next_base
WORKDIR /app/front
USER node

FROM next_base as next_prod
COPY --chown=node:node . /app/front
