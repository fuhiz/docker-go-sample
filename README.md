# docker-go-sample

Development environment for golang

## Build

```
docker-compose up -d --build
```

## Project setup

```
docker-compose exec web bash
```

```
go run cmd/migrate/main.go -exec up
```

```
go run cmd/server/main.go
```

## API usage

- Create user

```
curl localhost:8082/api/v1/users -X POST -H "Content-Type: application/json" -d '{"name": "test", "age":30}'
```

- Find User

```
curl localhost:8082/api/v1/users/:id
```

- User list

```
curl localhost:8082/api/v1/users
```

- Update user

```
curl localhost:8082/api/v1/users/:id -X PATCH -H "Content-Type: application/json" -d '{"name": "update", "age":31}'
```

- Delete user

```
curl localhost:8082/api/v1/users/:id -X DELETE
```

## Test

```
APP_MODE=test go run cmd/migrate/main.go -exec up
```

```
APP_MODE=test go test -v ./pkg/...
```

```
APP_MODE=test go run cmd/migrate/main.go -exec down
```
