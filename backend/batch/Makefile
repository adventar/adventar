.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/update_entry_site_info update_entry_site_info/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --stage prod --verbose

run: build
	sls invoke local -f ${f}
