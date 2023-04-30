setup:
	go get -u
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go install github.com/m-mizutani/zenv@latest

start:
	zenv go run .

sqlc-gen:
	sqlc generate
