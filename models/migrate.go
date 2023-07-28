package models

import "bugtracker/storage"

func Migrate() error {
	storage.ApplicationDB.AutoMigrate(&User{}, &Board{}, &Column{})
	return nil
}
