package containers

import (
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"log"
)

type RabbitMqContainer struct {
	pool      *dockertest.Pool
	resource  *dockertest.Resource
	imagename string
	opts      rabbitmq.Opts

	broker rabbitmq.IRabbitMq
}

type IRabbitMqContainer interface {
	Create(opts rabbitmq.Opts) error
	Connect()
	Flush()
}

func NewRabbitMqContainer(pool *dockertest.Pool) IRabbitMqContainer {
	return RabbitMqContainer{pool: pool, imagename: "rabbitmq-testcontainer"}
}

func (container RabbitMqContainer) Create(opts rabbitmq.Opts) error {
	dockerOpts := dockertest.RunOptions{
		Repository: "rabbitmq",
		Tag:        "3-management",
		Env: []string{
			"RABBITMQ_DEFAULT_USER=" + opts.Username,
			"RABBITMQ_DEFAULT_PASS=" + opts.Password,
		},
		ExposedPorts: []string{"5672", "15672"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5672":  {docker.PortBinding{HostIP: "0.0.0.0", HostPort: "5672"}},
			"15672": {docker.PortBinding{HostIP: "0.0.0.0", HostPort: "15672"}},
		},
		Name: container.imagename,
	}

	resource, err := container.pool.RunWithOptions(&dockerOpts)
	if err != nil {
		log.Fatalf("Could not start resource (RabbitMQ Test Container): %s", err.Error())
	}

	container.opts = opts
	container.resource = resource
	return nil
}

func (container RabbitMqContainer) Connect() {
	broker, err := rabbitmq.New(container.opts)
	if err != nil {
		panic(err)
	}

	container.broker = broker
}

func (container RabbitMqContainer) Flush() {}
