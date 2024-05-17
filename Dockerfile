#docker build -t web-server . && docker run -d -p 8080:8080 web-server

# Используем готовый образ Golang версии 1.21.1 как основу для нашего контейнера
FROM golang:1.21.1

# Устанавливаем рабочую директорию в корень контейнера
WORKDIR /
#usr/src/appScript

# Копируем все файлы из текущей директории в корневую директорию контейнера
COPY . /
#/usr/src/appScript

# mkdir -p /usr/src/appScript &&
ENV DATABASE_HOST=db

# Собираем ваш Go-проект (компилируем) с помощью команды go build и файлом main.go в корень контейнер
RUN apt-get update && \
    make build && \
    # go build -o main /cmd && \
    apt-get update && apt-get install -y postgresql-client nano 

# Указываем команду, которая будет выполняться при запуске контейнера
# Измеенить имя запускаюшего файла, например, appScript
CMD ["./main"]

# Открываем порт 8080, чтобы контейнер мог принимать входящие сетевые запросы на этом порту
EXPOSE 8080


# docker exec -it app /bin/bash
# apt update
# apt install netcat-openbsd
# install nmap-ncat
# psql -U user -d restapi_script && \dt \l

