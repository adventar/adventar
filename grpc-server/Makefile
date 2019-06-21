NAME = grpc-server

export GO111MODULE=on

proto:
	protoc -I ../protobuf -I $(GOPATH)/src/github.com/googleapis/googleapis ../protobuf/adventar/v1/*.proto --go_out plugins=grpc:.

build: proto
	go build -o $(NAME)

run: build
	./$(NAME)
