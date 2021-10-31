FROM golang:alpine

RUN mkdir /app
WORKDIR /app

CMD [ "go", "run", "/app/src/" ]