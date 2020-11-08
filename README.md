# fb.testcontainers

`Integration testing` is the phase in software testing in which individual software modules are combined and tested as a group ([wiki](https://en.wikipedia.org/wiki/Integration_testing)).
`Container testing` is on the other hand allows you to test your dockerized application end-to-end with 3rd party tools as if were in the production environment and without any dependencies.

So how we can implement it?

## Installation

To run a PostgreSql
```sh
$ docker run --name mypostgres -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=testcontainers -d -p 5432:5432 postgres
``` 

To run a RabbitMq;
```sh
$ docker run -d --hostname my-rabbit --name myrabbit -e RABBITMQ_DEFAULT_USER=guest -e RABBITMQ_DEFAULT_PASS=123456 -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```
and create a virtual host called as ``demand``.

## Tests

To run all tests;
```sh
$ go test -v ./...
```

When you run the integration tests and after the result, clear all the containers as below;
```sh 
$ docker rm -f $(docker ps -aq)
```

## Articles

- Integration Testing with Golang (Test Containers) - [Link](https://bit.ly/38k6THn "Integration Testing with Golang (Test Containers")