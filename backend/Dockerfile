# syntax=docker/dockerfile:1

FROM golang:1.17.5-alpine

WORKDIR /app

COPY ["go.mod", "go.sum", "*.go", "./"]
COPY .env.production .env

RUN go mod download

COPY src/ src/
COPY templates/ templates/

RUN go build -o /server

ENV GIN_MODE release
ENV NODE_ENV production

EXPOSE 8080
CMD ["/server"]