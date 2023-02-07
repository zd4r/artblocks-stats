# artblocks-stats

Small API to gather specific arblocks collection holders and their distribution by scores based on [Artacle API](https://artacle.github.io/api-docs/).

![build](https://github.com/zd4r/artblocks-stats/actions/workflows/main.yml/badge.svg)
[![codecov](https://codecov.io/gh/zd4r/artblocks-stats/branch/main/graph/badge.svg?token=5KTBZW0IH6)](https://codecov.io/gh/zd4r/artblocks-stats)
[![Go Report Card](https://goreportcard.com/badge/github.com/zd4rova/artblocks-stats)](https://goreportcard.com/report/github.com/zd4rova/artblocks-stats)
## Starting project
Just run:
```bash
$ make compose-build-up
```
And apply migrations with `docker usage` of [migrate](https://github.com/golang-migrate/migrate):
```bash
$ docker run -v migrations:/migrations --network host migrate/migrate -path=/migrations/ -database 'postgres://user:pass@localhost:5432/holders?sslmode=disable' up
```
Swagger API specification can be found at [http://localhost:8080/api-docs/](http://localhost:8080/api-docs/) (with default service port configuration).
## Tests
To start unit tests run:
```bash
$ make test
```
## Technical stack

- Backend building blocks
    - [labstack/echo/v4](https://github.com/labstack/echo)
    - [rabbitmq/amqp091-go](https://github.com/rabbitmq/amqp091-go)
    - [jmoiron/sqlx](github.com/jmoiron/sqlx)
        - [pq](github.com/lib/pq)
    - [golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate)
    - Utils
        - [google/wire](github.com/google/wire)
        - [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv)
        - [rs/zerolog](https://github.com/rs/zerolog)
- Infrastructure
    - Postgres
    - docker and docker-compose