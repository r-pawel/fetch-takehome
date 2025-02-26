FROM golang:1.24-alpine AS builder

WORKDIR /project/go-docker/

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o /project/go-docker/build/server cmd/server/main.go

FROM alpine:latest
COPY --from=builder /project/go-docker/build/server /project/go-docker/build/server

EXPOSE 8080
ENTRYPOINT [ "/project/go-docker/build/server" ]