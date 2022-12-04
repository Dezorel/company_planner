FROM golang:alpine

WORKDIR /app/company-schedule

COPY . /app/company-schedule

RUN go get "gopkg.in/yaml.v3"
RUN go get "github.com/go-sql-driver/mysql"

CMD go run main.go
