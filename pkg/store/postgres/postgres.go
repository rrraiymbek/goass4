package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Settings struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Conn struct {
	Pool *pgxpool.Pool
}

func New(settings Settings) (*Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		settings.User, settings.Password, settings.Host, settings.Port, settings.DBName)

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		Pool: pool,
	}

	return conn, nil
}

func (c *Conn) Stat() *pgxpool.Stat {
	return c.Pool.Stat()
}
