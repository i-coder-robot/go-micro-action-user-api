FROM bigrocs/golang-gcc:1.14 as builder

WORKDIR /Users/monster/GitHub/go-micro-action-user-api
COPY . .

ENV GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -a -installsuffix cgo -o bin/service