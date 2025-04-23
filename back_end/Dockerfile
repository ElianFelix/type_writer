# syntax=docker/dockerfile:1
ARG API_PORT=1323

FROM golang:1.23.8-alpine AS base

WORKDIR /app

COPY . .

RUN go mod download


FROM base AS api

RUN CGO_ENABLED=0 GOOS=linux go build -o /type_writer_api

EXPOSE ${API_PORT}

CMD ["/type_writer_api"]


FROM base AS migrate

# COPY --from=0 db/migration.go /

# COPY --from=base migrations /migrations

CMD go run db/migration.go

# RUN CGO_ENABLED=0 GOOS=linux go build -o /migrate migrate.go
