package main

import (
	"time"

	"github.com/johnbenedictyan/One-CV-Test/config"
	"github.com/johnbenedictyan/One-CV-Test/infra/database"
	"github.com/johnbenedictyan/One-CV-Test/infra/logger"
	"github.com/johnbenedictyan/One-CV-Test/routers"
	"github.com/spf13/viper"
)

func main() {
	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	masterDSN, replicaDSN, testDSN := config.DbConfiguration()

	if err := database.DBConnection(masterDSN, replicaDSN, testDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	router := routers.Routes()

	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
