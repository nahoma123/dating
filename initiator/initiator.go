package initiator

import (
	"context"
	"dating/internal/handler/middleware"
	"dating/platform/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func Initiate() {
	log := logger.New(InitLogger())
	log.Info(context.Background(), "logger initialized")

	log.Info(context.Background(), "initializing config")
	configName := "config"
	if name := os.Getenv("CONFIG_NAME"); name != "" {
		configName = name
		log.Info(context.Background(), fmt.Sprintf("config name is set to %s", configName))
	} else {
		log.Info(context.Background(), "using default config name 'config'")
	}
	InitConfig(configName, "config", log)
	log.Info(context.Background(), "config initialized")

	log.Info(context.Background(), "initializing database")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("database.url")))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	log.Info(context.Background(), "database initialized")

	log.Info(context.Background(), "initializing migration")
	InitiateMigration(viper.GetString("migration.path"), viper.GetString("database.url"), log)
	log.Info(context.Background(), "migration initialized")

	log.Info(context.Background(), "initializing persistence layer")
	CreateIndexes(log, client.Database(viper.GetString("database.name")))
	persistence := InitPersistence(client.Database(viper.GetString("database.name")), log)
	log.Info(context.Background(), "persistence layer initialized")

	log.Info(context.Background(), "initializing platform layer")
	platformLayer := InitPlatformLayer(log)
	log.Info(context.Background(), "platform layer initialized")

	log.Info(context.Background(), "initializing module")
	module := InitModule(persistence, viper.GetString("private_key"), platformLayer, log)
	log.Info(context.Background(), "module initialized")

	log.Info(context.Background(), "initializing handler")
	handler := InitHandler(module, log)
	log.Info(context.Background(), "handler initialized")

	log.Info(context.Background(), "initializing server")
	server := gin.New()
	server.Use(middleware.GinLogger(log))
	server.Use(ginzap.RecoveryWithZap(log.GetZapLogger().Named("gin.recovery"), true))
	server.Use(middleware.ErrorHandler())
	if viper.GetBool("dev") {
		server.Use(InitCORS())
	}
	log.Info(context.Background(), "server initialized")

	log.Info(context.Background(), "initializing metrics route")
	InitMetricsRoute(server, log)
	log.Info(context.Background(), "metrics route initialized")

	log.Info(context.Background(), "initializing router")
	v1 := server.Group("/v1")
	InitRouter(server, v1, handler, module, log, viper.GetString("public_key"))
	log.Info(context.Background(), "router initialized")

	srv := &http.Server{
		Addr:    viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		Handler: server,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	go func() {
		log.Info(context.Background(), "server started",
			zap.String("host", viper.GetString("server.host")),
			zap.Int("port", viper.GetInt("server.port")))
		log.Info(context.Background(), fmt.Sprintf("server stopped with error %v", srv.ListenAndServe()))
	}()
	sig := <-quit
	log.Info(context.Background(), fmt.Sprintf("server shutting down with signal %v", sig))
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("server.timeout"))
	defer cancel()

	log.Info(ctx, "shutting down server")
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("error while shutting down server: %v", err))
	} else {
		log.Info(context.Background(), "server shutdown complete")
	}
}
