FROM golang:1.15-alpine

RUN apk add --no-cache git

WORKDIR /app/authorizer

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .