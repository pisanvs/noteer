FROM node:alpine

WORKDIR /app

RUN yarn install

EXPOSE 3000

CMD [ "yarn", "dev" ]
