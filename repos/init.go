package repos

import "bugtracker/storage"

type repoType struct {
	BoardRepository         *boardRepo
	UserRepository          *userRepo
	RoleRepository          *roleRepo
	BoardUserRoleRepository *boardUserRoleRepo
}

var Repos *repoType

func InitializeRepositories() {
	Repos = &repoType{
		BoardRepository:         newBoardRepo(storage.ApplicationDB),
		UserRepository:          newUserRepo(storage.ApplicationDB),
		RoleRepository:          newRoleRepo(storage.ApplicationDB),
		BoardUserRoleRepository: newBoardUserRoleRepo(storage.ApplicationDB),
	}
}
