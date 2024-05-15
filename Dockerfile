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
# Собираем ваш Go-проект (компилируем) с помощью команды go build и файлом main.go в корень контейнер
RUN apt-get update && go build -o main /cmd && \
    apt-get update && apt-get install -y postgresql-client nano

# RUN  service postgresql start &&  psql -U user -d restapi_script -f migrations/*.up.sql 
    

# Указываем команду, которая будет выполняться при запуске контейнера
# Измеенить имя запускаюшего файла, например, appScript
CMD ["./main"]

# Открываем порт 8080, чтобы контейнер мог принимать входящие сетевые запросы на этом порту
EXPOSE 8080


# docker exec -it bbddd13c01bc /bin/bash
# apt update
# apt install netcat-openbsd
# install nmap-ncat
