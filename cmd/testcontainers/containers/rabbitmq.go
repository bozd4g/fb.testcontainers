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
}

type IRabbitMqContainer interface {
	C() RabbitMqContainer
	Create() error
	Connect() rabbitmq.IRabbitMq
	Flush(queues ...string) error
}

func NewRabbitMqContainer(pool *dockertest.Pool) IRabbitMqContainer {
	opts := rabbitmq.Opts{
		Username: "testcontainer",
		Password: "Aa123456.",
	}

	return RabbitMqContainer{pool: pool, opts: opts, imagename: "rabbitmq-testcontainer"}
}

func (container RabbitMqContainer) C() RabbitMqContainer {
	return container
}

func (container RabbitMqContainer) Create() error {
	if IsRunning(*container.pool, container.imagename) {
		return nil
	}

	dockerOpts := dockertest.RunOptions{
		Repository: "rabbitmq",
		Tag:        "3-management",
		Env: []string{
			"RABBITMQ_DEFAULT_USER=" + container.opts.Username,
			"RABBITMQ_DEFAULT_PASS=" + container.opts.Password,
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

	container.resource = resource
	return nil
}

func (container RabbitMqContainer) Connect() rabbitmq.IRabbitMq {
	var broker rabbitmq.IRabbitMq
	if err := container.pool.Retry(func() error {
		var err error
		broker, err = rabbitmq.New(container.opts)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return broker
}

func (container RabbitMqContainer) Flush(queues ...string) error {
	broker := container.Connect()

	var err error
	for _, queue := range queues {
		err = broker.Purge(queue)
		if err != nil {
			break
		}
	}

	return err
}