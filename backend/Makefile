setup:
	go get -u
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

start:
	zenv go run .

sqlc-gen:
	rm -rf pkg/gen/sqlc
	sqlc generate
