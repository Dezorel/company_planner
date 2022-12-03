FROM golang:alpine

WORKDIR /app/gateway

COPY . /app/gateway

RUN go get "gopkg.in/yaml.v3"

CMD go run main.go