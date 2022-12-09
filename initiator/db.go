package initiator

import (
	"context"
	"dating/platform/logger"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func InitDB(url string, log logger.Logger) *pgx.Conn {
	config, err := pgx.ParseConfig(url)
	config.Logger = log
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to connect to database: %v", err))
	}
	return conn
}
