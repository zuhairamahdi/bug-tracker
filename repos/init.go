package repos

import "bugtracker/storage"

type repoType struct {
	BoardRepository *boardRepo
}

var Repos *repoType

func InitializeRepositories() {
	Repos = &repoType{
		BoardRepository: NewBoardRepo(storage.ApplicationDB),
	}
}
