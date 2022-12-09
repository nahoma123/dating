package initiator

import (
	"context"
	"dating/platform/logger"
	"dating/platform/pgxadapter"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/jackc/pgx/v4"
)

func InitEnforcer(path string, conn *pgx.Conn, log logger.Logger) *casbin.Enforcer {
	adapter, err := pgxadapter.NewAdapterWithDB(conn)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to create adapter: %v", err))
	}

	enforcer, err := casbin.NewEnforcer(path, adapter)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to create enforcer: %v", err))
	}

	return enforcer
}
