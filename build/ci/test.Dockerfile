FROM golang:1.12

ENV CC_TEST_REPORTER_ID fd819ab4bfeccbd0a1376b4ed60774f3a716912aa7578672f7c08324c4d22745

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
RUN go get github.com/vektra/mockery/cmd/mockery
RUN curl -o cc-test-reporter -fSsL https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 && \
  chmod +x cc-test-reporter && \
  ./cc-test-reporter before-build

COPY . .

RUN go generate ./...

RUN go test -coverprofile c.out ./...
RUN ./cc-test-reporter after-build --exit-code $? --prefix github.com/emman27/gitignore
