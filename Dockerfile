FROM golang:1.25.5-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download && \
    go build -o /app/bin/auth_server ./cmd/server/main.go

FROM scratch

WORKDIR /root/
COPY --from=builder /app/bin/auth_server .
COPY --from=builder /app/env .

CMD ["./auth_server", "-config-path", ".env"]
