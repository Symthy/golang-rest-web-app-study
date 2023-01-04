# build
FROM golang:1.19.4-bullseye as deploy-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags  "-w -s" -o go-rest-app


# deply
FROM debian:bullseye-slim as deploy
RUN apt-get update
COPY --from=deploy-builder /app/go-rest-app .
CMD ["./go-rest-app"]


# dev (for hot reload)
FROM golang:1.19.4 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]