package repos

import "bugtracker/storage"

type repoType struct {
	BoardRepository *boardRepo
	UserRepository  *userRepo
}

var Repos *repoType

func InitializeRepositories() {
	Repos = &repoType{
		BoardRepository: newBoardRepo(storage.ApplicationDB),
		UserRepository:  newUserRepo(storage.ApplicationDB),
	}
}
