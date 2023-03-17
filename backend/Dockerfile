FROM golang:1.20.2 AS build
ARG BUILD_VERSION
ARG GITHUB_TOKEN
COPY . /app

WORKDIR /app
RUN go build -o adventar-backend

FROM gcr.io/distroless/base
COPY --from=build /app/adventar-backend /adventar-backend

WORKDIR /
EXPOSE 8080
ENTRYPOINT ["/adventar-backend"]
