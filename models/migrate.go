package models

import (
	"bugtracker/storage"
)

func Migrate() error {
	storage.ApplicationDB.AutoMigrate(&User{}, &Board{}, &Column{}, &Task{}, &Comment{})
	// MigrateInitialUser()
	return nil
}

// func MigrateInitialUser() {
// 	if err := storage.ApplicationDB.AutoMigrate(&User{}); err == nil && storage.ApplicationDB.Migrator().HasTable(&User{}) {
// 		if err := storage.ApplicationDB.First(&User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
// 			//Insert seed data
// 			admin := structs.NewUser{
// 				Id
// 				Username:  "admin",
// 				Email:     "admin@localhost",
// 				FirstName: "Admin",
// 				LastName:  "Admin",
// 				Password:  "default",
// 			}
// 			repos.Repos.UserRepository.Create(admin)
// 		}
// 	}

// }
