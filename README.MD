## Oukoud Backend

This is the backend for the Oukoud project. It is a GraphQL API built with [gqlgen](https://gqlgen.com/) and [gin web framework](), it uses [MongoDB](https://www.mongodb.com/) as the database, and Firebase for authentication.

### Technologies
List of the main technologies and tools used in the project:
- [Go 1.22](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [MongoDB](https://www.mongodb.com/)
- [gin web framework]()

### Installation
Instructions to run the project locally:

1. Clone the repository:
```bash
git clone git clone github.com/behrouz-rfa/kentech.git
cd kentech
```

2. Install the dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root of the project and add the following environment variables:
```bash
# Server Variables
SERVER_PORT=8080# default:80
SERVER_HOST=localhost# default:"localhost"
GIN_MODE=debug# default:"release"

# MongoDB Variables
DB_HOST=localhost
DB_PORT=27017
DB_NAME=oukoud
DB_USER=dbmaster
DB_PASS=password
DB_SSL=false
DB_CLUSTERED=false 

# Environment Variables
ENV=dev
LOG_LEVEL=debug


```

4. Run the project:
```bash
go run ./cmd/server/main.go
```
The server will be running on `http://localhost:8080`

### Support swagger
```
http://localhost:8080/swagger/index.html#/
```

### Docker
Instructions to run the project with docker:

1. Clone the repository:
```bash
git clone github.com/behrouz-rfa/kentech.git
cd kentech
```

2. Run the project:
```bash
docker-compose up
```

### GraphQL Playground
The GraphQL Playground is available at `http://localhost:8080`

### JustFile
The project uses [Just] to automate some tasks. To install it, check the [installation instructions](https://github.com/casey/just#installation).

### Justfile commands
List of the available commands in the Justfile:
- `just build [server]`: Build the project (server by default)
- `just run`: Run the project
- `just test [all | unit | e2e | it]`: Run the tests (all by default)
- `just lint`: Run the linter

### Tests
The project has three types of tests:
- Unit tests: Unit tests are located in the `internal` package and are mainly used to test the business logic.
- Integration tests: Integration tests are located in the `tests` package and are mainly used to test the API.
- End-to-end tests: End-to-end tests are located in the `e2e` package and are mainly used to test the API.

To run the tests:
```bash
just test
```


### Linter
The project uses [golangci-lint](https://golangci-lint.run/) to lint the project. To run the linter:
```bash
just lint
```
