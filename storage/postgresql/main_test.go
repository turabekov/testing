package postgresql

import (
	"app/config"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	categoryTestRepo *categoryRepo
	productTestRepo  *productRepo
	clientTestRepo   *clientRepo
	orderTestRepo    *orderRepo
)

func TestMain(m *testing.M) {
	cfg := config.Load()

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(pool)
	}

	categoryTestRepo = NewCategoryRepo(pool)
	productTestRepo = NewProductRepo(pool)
	clientTestRepo = NewClientRepo(pool)
	orderTestRepo = NewOrderRepo(pool)

	os.Exit(m.Run())
}
