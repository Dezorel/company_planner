FROM golang:alpine

WORKDIR /app/gateway

COPY . /app/gateway

RUN go get "gopkg.in/yaml.v3"
RUN go get "github.com/gorilla/mux"

CMD go run main.go