http-rest-api

Как запустить:

    Нужно скопировать docker-compose.yaml файл и в терминале запустить команды:

    Поднимаем сервис:

    docker compose up -d

    Запускаем миграции для создания таблицы юзеров:

    docker compose -f docker-compose.yaml --profile tools run --rm migrate up

    После можно удалить таблицу юзеровЖ

    docker compose -f docker-compose.yaml --profile tools run --rm migrate down

Или же скопировать проект и запустить комаду:

    make

    - запустятся тесты, сгенирируется бинарник, поднимутся контейнеры с базой и сервисом, создаться таблица юзеров, сервис готов к запросам.


Описание:

    Создание юзера:

    POST     http://localhost:8080/users

   Удаление юзера:

    DELETE     http://localhost:8080/users

   Вход юзера, создание сессии:

    POST  http://localhost:8080/users

    Получение авторизованного юзера:

    GET  http://localhost:8080/private/whoami
