APP_NAME=kbank

build:
	go build -o $(APP_NAME) ./cmd/server

run: build
	./$(APP_NAME)

generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/auth_service.proto