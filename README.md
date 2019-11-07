# twitter-clone-api

## Usage

```sh
$ git clone https://github.com/NasSilverBullet/twitter-clone-api.git
$ cd twitter-clone-api

$ docker-compose up
$ docker-compose run --rm db db/cli up
$ docker-compose run --rm db db/cli seed

$ curl -s localhost:8080/users | jq
[
  {
    "id": 1,
    "name": "Luke Skywalker",
    "email": "luke@example.com"
  },
  {
    "id": 2,
    "name": "Leia Organa",
    "email": "leia@example.com"
  },
  {
    "id": 3,
    "name": "Han Solo",
    "email": "han@example.com"
  },
  {
    "id": 4,
    "name": "Chewbacca",
    "email": "chewbacca@example.com"
  }
]
```

## Architecture

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
│   │   ├── router.go
│   │   └── sql_handler.go
│   ├── interfaces
│   │   ├── sql_handler.go
│   │   ├── user_handler.go
│   │   └── user_repository.go
│   ├── main.go
│   └── usecases
│       ├── user_interactor.go
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

## License

MIT License. See LICENSE.txt for more information.
