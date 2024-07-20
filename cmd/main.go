package main

import (
	"fmt"
	"product/api"
	"product/config"
	"product/pkg/db"
	log "product/pkg/logger"
	"product/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {

	cfg := config.Load()

	log := log.NewLogger(cfg.GeneralConfig)

	pgxdb, err := db.ConnectToDb(cfg.PgConfig)
	if err != nil {
		log.Error("error on connecting with database", logger.Error(err))
		return
	}

	log.Debug("successfully connected with database")

	fmt.Println(pgxdb)

	storage := storage.NewStorage(pgxdb, log)

	engine := api.Api(api.Options{
		Storage: storage,
		Log:     log,
	})

	log.Debug("server is running on", logger.String("port", cfg.GeneralConfig.HTTPPort))

	engine.Run(cfg.GeneralConfig.HTTPPort)
}
