FROM node:alpine

WORKDIR /app

RUN yarn

EXPOSE 3000

CMD [ "yarn", "dev" ]
