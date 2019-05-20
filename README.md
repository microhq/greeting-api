# Greeting API

The greeting api is used to create greetings

## Getting Started

### Run Micro API

```shell
micro api
```

### Run Greeting Service

```shell
go run greeting-srv/main.go
```

### Run Greeter API

```shell
go run main.go
```

### Curl API

```shell
curl -H 'Content-Type: application/json' -d '{"name": "Asim"}' http://localhost:8080/greeting/hello
```

Output

```
{
  "msg": "Hello Asim"
}
```
