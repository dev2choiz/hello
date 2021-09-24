FROM node:16-alpine as release
WORKDIR /app/front
#COPY --chown=node:node . /app/front
COPY . /app/front
RUN npm install
RUN npm run build
#ENV PORT 3000
#ENV HOST 0.0.0.0

#RUN chown -R node:node .next node_modules
#USER node

CMD ["npm", "run", "start"]
