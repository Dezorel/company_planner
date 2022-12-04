FROM golang:alpine

WORKDIR /app/cabinet-service

COPY . /app/cabinet-service

RUN go get "gopkg.in/yaml.v3"
RUN go get "github.com/go-sql-driver/mysql"

CMD go run main.go
