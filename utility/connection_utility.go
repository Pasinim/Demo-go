package utility

//in seguito le chiamate da fare nei test per connettersi

// usage:
// testDB := testhelpers.NewTestDatabase(t)
// defer testDB.Close(t)
// println(testDB.ConnectionString(t))

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestContainer struct {
	Instance testcontainers.Container
}

/*
*
Restituisce un container di Test
*/
func NewTestContainer() *TestContainer {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:alpine",
		ExposedPorts: []string{"5432/tcp"},
		AutoRemove:   true,
		Env: map[string]string{
			"POSTGRES_DB":       "demo",
			"POSTGRES_USER":     "demo",
			"POSTGRES_PASSWORD": "demo",
		},
		WaitingFor: wait.ForListeningPort("5432"),
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &TestContainer{
		Instance: postgres,
	}
}

func (db *TestContainer) Port() int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := db.Instance.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatal(err)
	}
	//require.NoError(t, err) serve se fa asser nil?
	return p.Int()
}

func (db *TestContainer) ConnectionString() string {
	return fmt.Sprintf("postgres://demo:demo@127.0.0.1:%d/demo?sslmode=disable", db.Port())
}

func (db *TestContainer) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, db.Instance.Terminate(ctx))
}
