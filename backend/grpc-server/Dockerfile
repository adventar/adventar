FROM golang:1.13.1-alpine3.10 as builder
RUN apk update && apk upgrade && apk add git gcc musl-dev
COPY . /app
WORKDIR /app
RUN go build -o grpc-server

FROM alpine:3.10
RUN apk update && apk upgrade && apk add ca-certificates
WORKDIR /go
COPY --from=builder /app/grpc-server ./
ENV TZ=Asia/Tokyo

CMD ["/go/grpc-server"]
