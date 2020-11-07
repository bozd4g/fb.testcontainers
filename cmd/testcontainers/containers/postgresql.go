package containers

import (
	"fmt"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/domain/user"
	"github.com/bozd4g/fb.testcontainers/pkg/postgresql"
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
	opts      postgresql.Opts
}

type IPostgreSqlContainer interface {
	C() PostgreSqlContainer
	Create() error
	Connect() *gorm.DB
	AutoMigrate(db *gorm.DB) error
	Flush(db *gorm.DB)
}

func NewPostgresqlContainer(pool *dockertest.Pool) IPostgreSqlContainer {
	opts := postgresql.Opts{
		Host:     "localhost",
		User:     "testcontainer",
		Password: "Aa123456.",
		Database: "testcontainer",
		Port:     5432,
	}

	return PostgreSqlContainer{pool: pool, opts: opts, imagename: "postgresql-testcontainer"}
}

func (container PostgreSqlContainer) C() PostgreSqlContainer {
	return container
}

func (container PostgreSqlContainer) Create() error {
	if IsRunning(*container.pool, container.imagename) {
		return nil
	}

	dockerOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_USER=" + container.opts.User,
			"POSTGRES_PASSWORD=" + container.opts.Password,
			"POSTGRES_DB=" + container.opts.Database,
		},
		ExposedPorts: []string{strconv.Itoa(container.opts.Port)},
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port(strconv.Itoa(container.opts.Port)): {{HostIP: "0.0.0.0", HostPort: strconv.Itoa(container.opts.Port)}},
		},
		Name: container.imagename,
	}

	resource, err := container.pool.RunWithOptions(&dockerOpts)
	if err != nil {
		log.Fatalf("Could not start resource (Postgresql Test Container): %s", err.Error())
		return err
	}

	container.resource = resource
	return nil
}

func (container PostgreSqlContainer) Connect() *gorm.DB {
	var db *gorm.DB
	if err := container.pool.Retry(func() error {
		defaultDsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
		dsn := fmt.Sprintf(defaultDsn, container.opts.Host, container.opts.User, container.opts.Password, container.opts.Database, container.opts.Port)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return db
}

func (container PostgreSqlContainer) AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(user.Entity{})
	if err != nil {
		return err
	}

	return nil
}

func (container PostgreSqlContainer) Flush(db *gorm.DB) {
	db.Exec("truncate table public.users")
}