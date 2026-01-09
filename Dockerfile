FROM golang:1.25.5-alpine AS builder

COPY . /app/auth
WORKDIR /app/auth

RUN go mod download && \
    go build -o /app/auth/bin/auth_server ./cmd/server/main.go

FROM alpine:latest

RUN apk add --no-cache postgresql-client

WORKDIR /root/
COPY --from=builder /app/auth/bin/auth_server .
COPY prod.env .
COPY wait-for-postgres.sh /root/wait-for-postgres.sh

RUN chmod +x /root/wait-for-postgres.sh

ENTRYPOINT ["/bin/sh", "-c", "./wait-for-postgres.sh && ./auth_server -config-path prod.env"]