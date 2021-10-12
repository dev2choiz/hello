FROM node:16-alpine as release
WORKDIR /app/front-next
#COPY --chown=node:node . /app/front-next
COPY . /app/front-next

RUN npm install
RUN npm run build

#USER node
#RUN chown -R node:node .next node_modules

CMD ["npm", "run", "start"]
