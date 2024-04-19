package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lushenle/plam/pkg/api"
	"github.com/lushenle/plam/pkg/db"
	"github.com/lushenle/plam/pkg/log"
	"github.com/lushenle/plam/pkg/token"
	"github.com/lushenle/plam/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//	@Title			PLAM API
//	@Description	This is a sample server for PLAM API.
//	@version		1.0

//	@termsOfService	http://localhost:8080/api/v1/terms/
//	@contact.name	API Support
//	@contact.url	http://localhost:8080/api/v1/support
//	@contact.email	lushenle@gmail.com

//	@host		localhost:8080
//	@BasePath	/v1

//	@schemes	http https
//	@produce	json
//	@consumes	json

//	@securityDefinitions.apiKey	apiKeyAuth
//	@in							header
//	@name						Authorization

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	// Setup config
	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	logLevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
	plugin, closer := log.NewFilePlugin(config.Logfile, logLevel[config.Loglevel])
	defer closer.Close()
	logger := log.NewLogger(plugin)
	logger.Info("service starting...")

	// Tokenmaker
	tokenMaker, err := token.NewPasetoMaker(config.Server.TokenSymmetricKey)
	if err != nil {
		logger.Fatal("cannot create token maker", zap.String("tokenMaker", err.Error()))
	}

	// Set database
	conn, err := pgxpool.New(context.Background(), config.Database.DataSourceName)
	if err != nil {
		logger.Fatal("db connection failed", zap.String("db", err.Error()))
	}

	// Database schema migration
	if err = util.DBMigration(config.Database.MigrationURL, config.Database.DataSourceName); err != nil {
		logger.Fatal("db migration failed", zap.String("db", err.Error()))
	}
	logger.Info("db migrated successfully")
	store := db.NewStore(conn)

	srv := api.NewServer(config, api.WithStore(store), api.WithLogger(logger), api.WithTokenMaker(tokenMaker))
	if err := srv.Start(config.Server.ServerAddress); err != nil {
		logger.Fatal("failed to run server", zap.String("server", err.Error()))
	}
}
