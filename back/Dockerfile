# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
COPY ../configs/config.json ./config.json

RUN go build -o ./sso ./src

ENV TZ=America/Argentina/Buenos_Aires

EXPOSE 3232

CMD [ "./sso" ]