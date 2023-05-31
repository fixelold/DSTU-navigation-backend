package postgresql

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"navigation/internal/config"
)

// Интерфейс клиента
type Client interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

// Создание нового клиента
func NewClient(ctx context.Context, cfg config.AppConfig) Client {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Storage.Username,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.Database,
	)
	// Подключение к базе данных
	conn, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("can't connect to database %s", err.Error())
	}
	return conn
}
