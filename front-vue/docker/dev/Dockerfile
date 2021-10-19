FROM node:16 as vue_base
WORKDIR /app/front-vue
RUN npm install -g npm@latest
RUN npm install -g @vue/cli

USER node

FROM next_base as vue_prod
COPY --chown=node:node . /app/front-vue
