#  API-приложение
##
## Обзор
API-приложение предоставляет RESTful API для запуска bash-скриптов. Приложение позволяет пользователям создавать, получать и удалять скрипты, а также выполнять их и сохранять результаты в базе данных. Преимущественно использовались стандартные библиотеки Go для реализации функциональности.

##
## Функциональность:

API-приложение предоставляет следующую функциональность:
- Создание скрипта: Создает новый скрипт, выполняет переданную bash-команду и сохраняет результат в базе данных.
- Получение списка скриптов: Получает список всех скриптов.
- Получение скрипта по ID: Получает один скрипт по его ID.
- Удаление скрипта по ID: Удаляет один скрипт по его ID.

##
## Технологический стек

- Язык программирования: Go 1.21.1
- База данных: Postgres
- Операционная система: MacOS

##
## Настройка базы данных
Для запуска приложения и тестов нужно создать следующие базы данных: - restapi_script, - restapi_test
- make db (используя Makefile)


##
## Запуск сервера
 
- go run cmd/main.go
- make (используя Makefile)
- docker-compose up --build (используя Docker)

##
## Запуск тестов
- cd tests && go test
- make tests (используя Makefile)

