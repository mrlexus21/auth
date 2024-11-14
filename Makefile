LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user/v1
	protoc --proto_path api/user/v1 \
	--go_out=pkg/user/v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go.exe \
	--go-grpc_out=pkg/user/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc.exe \
	api/user/v1/user.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/server/main.go

copy-to-server:
	scp service_linux root@188.130.207.122:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.yandex/crpa9fe1mvd86ibqm7qn/test-server:v0.0.1 .
	#docker login -u mrlexus21 -p CRgAAAAA6ELvdoGnA7EL4qSbkICKeDkVldDC2OOU cr.selcloud.ru/olezhek28
	echo y0_AgAAAAADQna9AATuwQAAAAEYZvYPAADxQEN56WxMCpDw7t3NqCQr8ikZNw|docker login --username oauth --password-stdin cr.yandex
	docker push cr.yandex/crpa9fe1mvd86ibqm7qn/test-server:v0.0.1