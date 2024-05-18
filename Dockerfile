# Используем готовый образ Golang версии 1.21.1 как основу для нашего контейнера
FROM golang:1.21.1

# Определяем среду окружения
ENV env=dev
ENV DATABASE_HOST=db

# Устанавливаем рабочую директорию в корень контейнера
WORKDIR /appScript

# Копируем все файлы из текущей директории в корневую директорию контейнера
COPY . /appScript

# Собираем ваш Go-проект (компилируем) с помощью команды go build и файлом main.go в корень контейнер
RUN make build && \
#    go build -o main /appScript/cmd  && \
    apt-get update && \
    apt-get install iputils-ping  -y postgresql-client nano  netcat-openbsd 
   
RUN rm -rf /root/.cache/go-build/*

# Указываем команду, которая будет выполняться при запуске контейнера
CMD ["./main"]

# Открываем порт 8080, чтобы контейнер мог принимать входящие сетевые запросы на этом порту
EXPOSE 8080


