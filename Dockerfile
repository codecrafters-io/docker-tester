FROM golang:1.21-alpine

RUN apk add --no-cache --upgrade 'curl>=8.5'

# Download docker-explorer
ARG docker_explorer_version=v18
RUN curl --fail -Lo /usr/local/bin/docker-explorer https://github.com/codecrafters-io/docker-explorer/releases/download/${docker_explorer_version}/${docker_explorer_version}_linux_amd64
RUN chmod +x /usr/local/bin/docker-explorer

# Development deps
RUN apk add --no-cache --upgrade 'make>=4.4' 'gcc>=12.2' 'musl-dev>=1.2' 'git>=2.40'

WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

# Cache go modules
RUN go mod download

COPY . /app
