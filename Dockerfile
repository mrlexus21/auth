FROM golang:1.23-alpine AS builder

COPY . /github.com/mrlexus21/auth/source/
WORKDIR /github.com/mrlexus21/auth/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/mrlexus21/auth/source/bin/auth_server .
COPY --from=builder /github.com/mrlexus21/auth/source/prod.env .

CMD ["./auth_server", "-config-path", "prod.env"]