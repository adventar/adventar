NAME = grpc-server

export GO111MODULE=on

protoc:
	protoc -I ../protobuf ../protobuf/adventar/v1/*.proto --go_out plugins=grpc:.

build: protoc
	go build -o $(NAME)

run: build
	./$(NAME)
