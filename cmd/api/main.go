// Package main is the entry point for the API server.
package main

import (
	"github.com/boymeetsblockchain/e-commerce/internal/config"
	"github.com/boymeetsblockchain/e-commerce/internal/database"
	"github.com/boymeetsblockchain/e-commerce/internal/logger"
	"github.com/gin-gonic/gin"
)

func main() {

	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	mainDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get database connection")
	}

	defer func() {
		if err := mainDB.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close database")
		}
	}()
	gin.SetMode(cfg.Server.GinMode)

	log.Info().Msg("starting server")

}
