NAME = update_entry_site_info

export GO111MODULE=on

build:
	go build -o $(NAME)

run: build
	./$(NAME)

test:
	go test ./...

docker-build:
	docker build -t hokaccha/adventar-update_entry_site_info .

deploy: docker-build
	docker push hokaccha/adventar-update_entry_site_info
