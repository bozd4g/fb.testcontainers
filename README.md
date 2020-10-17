# fb.testcontainers

`Integration testing` is the phase in software testing in which individual software modules are combined and tested as a group ([wiki](https://en.wikipedia.org/wiki/Integration_testing)).
`Container testing` is on the other hand allows you to test your dockerized application end-to-end with 3rd party tools as if were in the production environment and without any dependencies.

So how we can implement it?

## Installation 

To run a RabbitMq;
```
docker run -d --hostname my-rabbit --name myrabbit -e RABBITMQ_DEFAULT_USER=guest -e RABBITMQ_DEFAULT_PASS=123456 -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```
and create a virtual host called as ``demand``.

## Articles [WIP]

...