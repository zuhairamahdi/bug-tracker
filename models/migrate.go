package models

import "bugtracker/storage"

func Migrate() error {
	storage.ApplicationDB.AutoMigrate(&User{})
	return nil
}
