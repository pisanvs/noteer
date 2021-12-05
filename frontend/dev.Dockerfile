FROM node:16-alpine

WORKDIR /app

RUN yarn

EXPOSE 3000

ENTRYPOINT [ "yarn", "dev" ]
