# API по созданию сокращённых ссылок

## Описание
Этот сервис позволяет создавать сокращённые ссылки с возможностью выбора хранилища.

## Выбор хранилища
Тип хранилища задаётся через переменную окружения:
- `STORAGE_TYPE=postgres` - использовать PostgreSQL
- `STORAGE_TYPE=inmem` - использовать In-Memory хранилище

## Запуск
Для запуска сервиса используйте команду:
```sh
docker-compose up
