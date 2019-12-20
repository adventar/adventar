FROM golang:1.13.1-alpine3.10 as builder
RUN apk update && apk upgrade && apk add git gcc musl-dev tzdata
COPY . /app
WORKDIR /app
RUN go build -o update_entry_site_info

FROM alpine:3.10
RUN apk update && apk upgrade && apk add ca-certificates
WORKDIR /go
COPY --from=builder /app/update_entry_site_info ./
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
ENV TZ=Asia/Tokyo
CMD ["/go/update_entry_site_info"]
