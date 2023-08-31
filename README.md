# Сервис динамического сегментирования пользователей 

Микросервис для работы с сегментами пользователей, добавлением сегментов, их удалением, 
добавлением и удалением сегментов у пользователя, выводом актуальных сегментов пользователя

Используемые технологии:
- PostgreSQL (в качестве хранилища данных)
- Docker (для поднятия сервиса и БД в контейнерах)
- Swagger (для документирования API)
- golang-migrate/migrate (для миграций БД)
- pq (драйвер для работы с PostgreSQL)
- Gin (веб-фреймворк)
- Cleanenv (для работы с конфигом и переменными окружения)

## Особенности проекта
### Архитектура проекта:
Я постарался выстроить проект на основе принципов 
чистой архитектуры и [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
для того, чтобы в дальнейшем иметь возможность фунцкционально расширять данный сервис

### Работа с проектом
Перед началом работы необходимо в корне проекта создать файл .env и заполнить необходимые переменные окружения. Список этих переменных лежит в файле .env.example

Все действия с проектом происходят с помощью Makefile


### **Работа с сервисом** 
Поднять сервис локально
```
make run
```
Поднять сервис в контейнере
```
make up
```
Завершить работу сервиса
```
make down 
```
Пересобрать проект и поднять его
```
make docker-rebuild
```

### **Работа с миграциями**
Создать файлы миграций
```
make migration {file-name}
```
Накатить миграцию
```
make migrate
```
Откатить миграцию
```
make migrate-down
```
Дропнуть базу
```
make migrate-drop
```

## Swagger
Для работы с нижеперечисленными методами предоставляется интерфейс Swagger UI, который доступен по адресу:
```
http://localhost:8080/swagger/index.html
```` 

![Swagger](https://imageupload.io/ib/THElLxWDwBPnDsg_1693503262.png)

## Примеры запросов

### Метод создания сегмента
```curl
curl --location 'localhost:8080/api/v1/segments' \
--header 'Content-Type: application/json' \
--data '{
    "slug" : "test1"
}'
```
### Пример ответа
```json
{
    "id": 1
}
```
### Метод создания сегмента c заданным процентом пользователей 
```curl
curl --location 'localhost:8080/api/v1/segmentsProb' \
--header 'Content-Type: application/json' \
--data '{
    "slug" : "test1",
    "probability" : 0.1
}'
```
### Пример ответа
```json
{
    "id": 1
}
```
### Метод удаления сегмента
```curl
curl --location --request DELETE 'localhost:8080/api/v1/segments/test4'
```
### Пример ответа
```json
{
  "id": 4
}
```
### Метод добавления пользователя в сегмент
```curl
curl --location 'localhost:8080/api/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "vp.ruzanov@gmail.com",
    "segmentsToAdd" : ["test1", "test2", "test3"],
    "segmentsToDelete" : ["test4"]
}'
```
### Пример ответа
```json
{
  "id": 1
}
```
### Метод получения активных сегментов пользователя
```curl
curl --location 'localhost:8080/api/v1/users/1/segments'
```
### Пример ответа
```json
{
  "id": 1,
  "segments": [
    {
      "id": 1,
      "slug": "test1"
    },
    {
      "id": 2,
      "slug": "test2"
    },
    {
      "id": 3,
      "slug": "test3"
    }
  ]
}
```


