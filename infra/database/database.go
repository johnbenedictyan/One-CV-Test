package database

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var (
	DB                 *gorm.DB
	err                error
	DBErr              error
	testDB             *gorm.DB
	testDbErr          error
	testDbMigrationErr error
)

// DBConnection create database connection
func DBConnection(masterDSN, replicaDSN, testDSN string) error {
	var db = DB
	var testDb = testDB

	logMode := viper.GetBool("DB_LOG_MODE")
	debug := viper.GetBool("DEBUG")

	loglevel := logger.Silent
	if logMode {
		loglevel = logger.Info
	}

	db, err = gorm.Open(postgres.Open(masterDSN), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})

	testDb, testDbErr = gorm.Open(postgres.Open(testDSN), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})

	if !debug {
		db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{
				postgres.Open(replicaDSN),
			},
			Policy: dbresolver.RandomPolicy{},
		}))
	}

	if err != nil {
		DBErr = err
		log.Println("Db connection error")
		return err
	}

	if testDbErr != nil {
		DBErr = err
		log.Println("Test Db connection error")
		return err
	}

	err = db.AutoMigrate(migrationModels...)

	testDbMigrationErr = testDb.AutoMigrate(migrationModels...)

	if err != nil {
		return err
	}
	DB = db

	if testDbMigrationErr != nil {
		return testDbMigrationErr
	}
	testDB = testDb

	return nil
}

// GetDB connection
func GetDB() *gorm.DB {
	return DB
}

// GetDBError connection error
func GetDBError() error {
	return DBErr
}

// GetTestDB connection
func GetTestDB() *gorm.DB {
	return testDB
}

// GetTestDBError connection error
func GetTestDBError() error {
	return testDbErr
}
