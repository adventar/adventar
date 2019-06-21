FROM golang:1.12.5-alpine3.9 as builder
RUN apk update && apk upgrade && apk add git gcc musl-dev
COPY . /app
WORKDIR /app
RUN GO111MODULE=on go build -o grpc-server

FROM alpine:3.9
RUN apk update && apk upgrade && apk add ca-certificates
WORKDIR /go
COPY --from=builder /app/grpc-server ./
COPY service_account.json ./

CMD ["/go/grpc-server"]
