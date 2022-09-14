package datasource

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var pgpool *pgxpool.Pool

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "9.6",
		Env: []string{
			"POSTGRES_DB=postgres",
			"POSTGRES_PASSWORD=password",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		pgpool, err = pgxpool.Connect(context.Background(), fmt.Sprintf("postgresql://postgres:password@localhost:%s/postgres?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		return pgpool.Ping(context.Background())
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	mig, err := migrate.New("file://../../database/migration", fmt.Sprintf("postgresql://postgres:password@localhost:%s/postgres?sslmode=disable", resource.GetPort("5432/tcp")))
	if err != nil {
		log.Fatalln(err)
	}

	if err := mig.Up(); err != nil {
		log.Fatalln(err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestPostgres_GetAllUsers(t *testing.T) {
	p := NewPostgres(pgpool)
	users, err := p.GetAllUsers()
	if err != nil {
		t.Errorf("getalluser err=%s; want nil", err)
	}
	// not great for parallel tests.
	want := 500
	got := len(users)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
func TestPostgres_GetUserByID_success(t *testing.T) {
	p := NewPostgres(pgpool)
	user, err := p.GetUserByID(1)
	if err != nil {
		t.Errorf("getuserbyid(1) err=%s; want nil", err)
	}
	want := 1
	got := user.ID
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPostgres_GetUserByID_failure(t *testing.T) {
	p := NewPostgres(pgpool)
	user, err := p.GetUserByID(-1) // this should return nil user
	assert.Error(t, err, "should return error")
	assert.Nil(t, user, "user should be nil")
}
