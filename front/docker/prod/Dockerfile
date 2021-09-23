FROM node:16-alpine as release
WORKDIR /app/front
COPY --chown=node:node . /app/front
RUN npm install
RUN npm run build
USER node
