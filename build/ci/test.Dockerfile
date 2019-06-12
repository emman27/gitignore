FROM golang:1.12

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go get github.com/vektra/mockery/cmd/mockery
RUN go generate ./...

RUN go test ./...
