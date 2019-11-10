.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o image-server main.go

clean:
	rm -f image-server

deploy: clean build
	sls deploy --verbose
