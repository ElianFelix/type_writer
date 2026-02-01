# syntax=docker/dockerfile:1

FROM golang:1.23.8-alpine AS base

WORKDIR /app

COPY . .

RUN go mod download


FROM base AS api

RUN CGO_ENABLED=0 GOOS=linux go build -o /type_writer_api

CMD ["/type_writer_api"]


FROM base AS migrate

CMD go run db/migration.go
