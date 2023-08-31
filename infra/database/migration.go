package database

import (
	"github.com/johnbenedictyan/One-CV-Test/models"
)

// Add list of model add for migrations
// var migrationModels = []interface{}{&ex_models.Example{}, &model.Example{}, &model.Address{})}
var migrationModels = []interface{}{
	&models.Teacher{},
	&models.Student{},
	&models.Notification{}}
