# twitter-clone-api

## Usage

```sh
$ git clone https://github.com/NasSilverBullet/twitter-clone-api.git
$ cd twitter-clone-api

$ docker-compose up
$ docker-compose run --rm db db/cli up
$ docker-compose run --rm db db/cli seed

$ curl -s localhost:8080/users/1 | jq -r
{
  "id": 1,
  "name": "Luke Skywalker",
  "email": "luke@example.com",
  "created": "2019-11-09T17:51:01+09:00",
  "updated_at": "2019-11-09T17:51:01+09:00",
  "deleted_at": "0001-01-01T00:00:00Z"
}
```

## Architecture

![The Clean Architecture](
https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg
"The Clean Architecture"
)

Refs: <https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html>

```sh
$ tree -a -I "\.DS_Store|\.git"
.
├── .env
├── .gitignore
├── .realize.yaml
├── LICENSE.txt
├── README.md
├── app
│   ├── entities
│   │   └── user.go
│   ├── frameworks
│   │   ├── env.go
│   │   ├── logger.go
│   │   ├── router.go
│   │   ├── sql_handler.go
│   │   └── validator.go
│   ├── interfaces
│   │   ├── logger.go
│   │   ├── sql_handler.go
│   │   ├── user_handler.go
│   │   ├── user_repository.go
│   │   ├── user_repository_test.go
│   │   └── validator.go
│   ├── main.go
│   └── usecases
│       ├── user_interactor.go
│       ├── user_interactor_test.go
│       └── user_repository.go
├── db
│   ├── cli
│   ├── migrations
│   │   ├── 20191106231421_create_users_table.down.sql
│   │   └── 20191106231421_create_users_table.up.sql
│   └── seeds
│       └── users_table.sql
├── docker
│   ├── app
│   │   └── Dockerfile
│   └── db
│       ├── Dockerfile
│       └── my.cnf
├── docker-compose.yml
├── go.mod
└── go.sum
```

## HTTP methods

| Method | HTTP Mapping     | HTTP Request Body | HTTP Respose Body |
|--------|------------------|-------------------|-------------------|
| List   | GET  /users      | N/A               | Users             |
| Get    | GET  /users/{id} | N/A               | User              |
| Create | POST /users      | User              | UserID            |

## CURL exmaple

```sh
# GET /users
$ curl localhost:8080/users

# GET /users/{id}
$ curl localhost:8080/users/1

# POST /users
$ curl localhost:8080/users \
-H 'Content-Type: application/json' \
-d '{"name":"Obi-Wan Kenobi","email":"obi-wan@example.com"}'
```

## License

MIT License. See LICENSE.txt for more information.
