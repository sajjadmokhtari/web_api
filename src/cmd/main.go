package main

import (
	"GOLANG_CLEAN_WEB_API/src/api"
	"GOLANG_CLEAN_WEB_API/src/cache"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization





func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)

}
