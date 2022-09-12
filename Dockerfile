# syntax=docker/dockerfile:1

FROM golang:1.19-alpine3.15

RUN apk add build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN apk add postgresql-client

RUN go build -o gotstock-api

EXPOSE 8080

