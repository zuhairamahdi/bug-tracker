package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance
var ApplicationDB *gorm.DB

func InitializeDB() {
	DNS := "host=localhost user=postgres password=password dbname=bugtracker port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Logger.LogMode(logger.Info)
	Database = DbInstance{
		Db: db,
	}
	ApplicationDB = Database.Db
}
