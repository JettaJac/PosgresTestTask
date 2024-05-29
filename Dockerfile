# Используем готовый образ Golang версии 1.21.1 как основу для нашего контейнера
FROM golang:1.21.1 AS builder 

# Устанавливаем рабочую директорию в корень контейнера
WORKDIR /appScript

# Копируем все файлы из текущей директории в корневую директорию контейнера
COPY . /appScript

# Компилируем приложение специально для дальнейшего запуска в Alpine      
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /appScript/cmd  && \
    apt-get update && apt-get install iputils-ping  -y postgresql-client 

RUN rm -rf /root/.cache/go-build/*


# Создаем финальный Докер-образ на базе легковесного alpine
FROM alpine:latest                     

# Определяем среду окружения
ENV env=dev
ENV DATABASE_HOST=db

WORKDIR /root/
# Копируем собранное приложение из образа biulder
COPY --from=builder ./appScript .             

# Устанавливаем недостающие утилы
RUN apk add sudo && apk add bash && apk add make && apk add go

# Открываем порт 8080, чтобы контейнер мог принимать входящие сетевые запросы на этом порту
EXPOSE 8080

# Указываем команду, которая будет выполняться при запуске контейнера
CMD ["./main"]

