package containers

import (
	"fmt"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/app"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type PostgreSqlContainer struct {
	pool      *dockertest.Pool
	resource  *dockertest.Resource
	imagename string
	opts      app.PostgreSqlOpts

	db *gorm.DB
}

type IPostgreSqlContainer interface {
	Create(opts app.PostgreSqlOpts) error
	Connect()
	Flush()
}

func NewPostgresqlContainer(pool *dockertest.Pool) IPostgreSqlContainer {
	return PostgreSqlContainer{pool: pool, imagename: "postgresql-testcontainer"}
}

func (container PostgreSqlContainer) Create(opts app.PostgreSqlOpts) error {
	port := docker.Port(strconv.Itoa(opts.Port))
	dockerOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.3",
		Env: []string{
			"POSTGRES_USER=" + opts.User,
			"POSTGRES_PASSWORD=" + opts.Password,
			"POSTGRES_DB=" + opts.Database,
		},
		ExposedPorts: []string{strconv.Itoa(opts.Port)},
		PortBindings: map[docker.Port][]docker.PortBinding{
			port: {{HostIP: "0.0.0.0", HostPort: strconv.Itoa(opts.Port)}},
		},
		Name: container.imagename,
	}

	resource, err := container.pool.RunWithOptions(&dockerOpts)
	if err != nil {
		log.Fatalf("Could not start resource (Postgresql Test Container): %s", err.Error())
		return err
	}

	container.opts = opts
	container.resource = resource
	return nil
}

func (container PostgreSqlContainer) Connect() {
	defaultDsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	dsn := fmt.Sprintf(defaultDsn, container.opts.Host, container.opts.User, container.opts.Password, container.opts.Database, container.opts.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	container.db = db
}

func (container PostgreSqlContainer) Flush() {}
