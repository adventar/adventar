FROM golang:1.20.2 AS build
COPY . /app

WORKDIR /app
RUN go build -o adventar-backend

FROM gcr.io/distroless/base
COPY --from=build /app/adventar-backend /adventar-backend

ENV TZ=Asia/Tokyo
WORKDIR /
EXPOSE 8080
ENTRYPOINT ["/adventar-backend"]
