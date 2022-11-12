FROM golang:alpine

#путь к папке внутри образа
WORKDIR /app/company-service

#путь откуда - куда скопировать
COPY . /app/company-service

# подкачиваем необходимые пакеты
RUN go get "gopkg.in/yaml.v3"
RUN go get "github.com/go-sql-driver/mysql"

#запускаем приложение
CMD go run main.go
