FROM golang:1.17-alpine

RUN apk add curl

# Download docker-explorer
ARG docker_explorer_version=v18
RUN curl --fail -Lo /usr/local/bin/docker-explorer https://github.com/codecrafters-io/docker-explorer/releases/download/${docker_explorer_version}/${docker_explorer_version}_linux_amd64
RUN chmod +x /usr/local/bin/docker-explorer

# Development deps
RUN apk add make gcc musl-dev git

WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

# Cache go modules
RUN go mod download

COPY . /app
