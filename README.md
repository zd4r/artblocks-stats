# artblocks-stats
![build](https://github.com/zd4r/artblocks-stats/actions/workflows/main.yml/badge.svg)
[![codecov](https://codecov.io/gh/zd4r/artblocks-stats/branch/main/graph/badge.svg?token=5KTBZW0IH6)](https://codecov.io/gh/zd4r/artblocks-stats)
[![Go Report Card](https://goreportcard.com/badge/github.com/zd4rova/artblocks-stats)](https://goreportcard.com/report/github.com/zd4rova/artblocks-stats)

Small API to gather specific arblocks collection holders and their distribution by scores based on [Artacle API](https://artacle.github.io/api-docs/).

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
## Usage example
Getting collection (№399) holders distribution:
```bash
$ curl http://localhost:8080/v1/collections/399/stats | jq .

{
  "collection": {
    "id": 399,
    "holders_count": 220,
    "holders_distribution": {
      "by_commitment_score": {
        "[3 - 3.5)": 43,
        "[3.5 - 4)": 58,
        "[4 - 4.5)": 59,
        "[4 - 4.5]": 60
      },
      "by_portfolio_score": {
        "[3 - 3.5)": 75,
        "[3.5 - 4)": 78,
        "[4 - 4.5)": 50,
        "[4 - 4.5]": 17
      },
      "by_trading_score": {
        "[3 - 3.5)": 42,
        "[3.5 - 4)": 27,
        "[4 - 4.5)": 106,
        "[4 - 4.5]": 45
      }
    }
  }
}
```
Getting collection (№399) holders with scores:
```bash
$ curl http://localhost:8080/v1/collections/399/holders | jq .

{
  "holders": [
    {
      "address": "0xb9eb79e3e735ee636255dd8d65872a1287744e33",
      "tokens_amount": 20,
      "commitment_score": 3.798311122595867,
      "portfolio_score": 4.03269501537506,
      "trading_score": 4.494485581716837
    },
    {
      "address": "0xcbbea7ec33d60db283ab79bdac9ffbfa46a83134",
      "tokens_amount": 11,
      "commitment_score": 4.331444471045636,
      "portfolio_score": 3.6309881414241354,
      "trading_score": 4.247808248625762
    },
    
    ...
    
    {
      "address": "0xf9b7d79932b16c6bf8d08dbce15cd5e6942dd18f",
      "tokens_amount": 1,
      "commitment_score": 3,
      "portfolio_score": 3.364766609678407,
      "trading_score": 3
    }
  ]
}
```

## Technical stack
- Backend building blocks
    - [labstack/echo/v4](https://github.com/labstack/echo)
    - [jmoiron/sqlx](github.com/jmoiron/sqlx)
        - [pq](github.com/lib/pq)
    - [golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate)
    - Utils
        - [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv)
        - [rs/zerolog](https://github.com/rs/zerolog)
- Infrastructure
    - Postgres
    - docker and docker-compose