FROM golang:1.19.0

#путь к папке внутри образа
WORKDIR /app

#путь откуда - куда скопировать
COPY . /app

# подкачиваем необходимые пакеты
RUN go get "gopkg.in/yaml.v3

#запускаем приложение
CMD go run main.go
