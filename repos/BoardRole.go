package repos

type boardRoleRepo struct {
	storage *gorm.DB
}

func newBoardRoleRepo(storage *gorm.DB) *boardRoleRepo {
	
	return &boardRoleRepo{
		storage: storage,
	}
}

func (r *boardRoleRepo) GetAll() ([]models.BoardRole, error) {
	boardRoles := []models.BoardRole{}
	err := r.storage.Find(&boardRoles).Error
	return boardRoles, err
}

func (r *boardRoleRepo) Create(boardRole structs.BoardRole) (models.BoardRole, error) {
	newBoardRole := models.BoardRole{
		ID:         ulid.Make().String(),
		Title:      boardRole.Title,
		Desciption: boardRole.Description,
	}
	newBoardRole.Users = []models.User{}
	for _, user := range boardRole.Users {
		newBoardRole.Users = append(newBoardRole.Users, models.User{
			ID: user.ID,
		})
	if query := r.storage.Create(&newBoardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	return newBoardRole, nil
}
func (r *boardRoleRepo) Get(id string) (models.BoardRole, error)
{
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
}
func (r *boardRoleRepo) Update(id string, boardRole structs.BoardRole) (models.BoardRole, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	boardRole.Title = boardRole.Title
	boardRole.Description = boardRole.Description
	if query := r.storage.Save(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	return boardRole, nil
}
func (r *boardRoleRepo) Delete(id string) error {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return query.Error
	}
	if query := r.storage.Delete(&boardRole); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *boardRoleRepo) AddUser(id string, user structs.User) (models.BoardRole, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	boardRole.Users = append(boardRole.Users, models.User{
		ID: user.ID,
	})
	if query := r.storage.Save(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	return boardRole, nil
}

func (r *boardRoleRepo) RemoveUser(id string, user structs.User) (models.BoardRole, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	for i, u := range boardRole.Users {
		if u.ID == user.ID {
			boardRole.Users = append(boardRole.Users[:i], boardRole.Users[i+1:]...)
		}
	}
	if query := r.storage.Save(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	return boardRole, nil
}

func (r *boardRoleRepo) CheckIfHaveAccess(boardId string, userId string) (bool, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("board_id = ?", boardId).First(&boardRole); query.Error != nil {
		return false, query.Error
	}
	for _, user := range boardRole.Users {
		if user.ID == userId {
			return true, nil
		}
	}
	return false, nil
}

func (r *boardRoleRepo) GetUsers(boardRoleId string) ([]models.User, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", boardRoleId).First(&boardRole); query.Error != nil {
		return []models.User{}, query.Error
	}
	return boardRole.Users, nil
}package repos

type boardRoleRepo struct {
	storage *gorm.DB
}

func newBoardRoleRepo(storage *gorm.DB) *boardRoleRepo {
	
	return &boardRoleRepo{
		storage: storage,
	}
}

func (r *boardRoleRepo) GetAll() ([]models.BoardRole, error) {
	boardRoles := []models.BoardRole{}
	err := r.storage.Find(&boardRoles).Error
	return boardRoles, err
}

func (r *boardRoleRepo) Create(boardRole structs.BoardRole) (models.BoardRole, error) {
	newBoardRole := models.BoardRole{
		ID:         ulid.Make().String(),
		Title:      boardRole.Title,
		Desciption: boardRole.Description,
	}
	newBoardRole.Users = []models.User{}
	for _, user := range boardRole.Users {
		newBoardRole.Users = append(newBoardRole.Users, models.User{
			ID: user.ID,
		})
	if query := r.storage.Create(&newBoardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	return newBoardRole, nil
}
func (r *boardRoleRepo) Get(id string) (models.BoardRole, error)
{
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
}
func (r *boardRoleRepo) Update(id string, boardRole structs.BoardRole) (models.BoardRole, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	boardRole.Title = boardRole.Title
	boardRole.Description = boardRole.Description
	if query := r.storage.Save(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	return boardRole, nil
}
func (r *boardRoleRepo) Delete(id string) error {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return query.Error
	}
	if query := r.storage.Delete(&boardRole); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *boardRoleRepo) AddUser(id string, user structs.User) (models.BoardRole, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	boardRole.Users = append(boardRole.Users, models.User{
		ID: user.ID,
	})
	if query := r.storage.Save(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	return boardRole, nil
}

func (r *boardRoleRepo) RemoveUser(id string, user structs.User) (models.BoardRole, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", id).First(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	for i, u := range boardRole.Users {
		if u.ID == user.ID {
			boardRole.Users = append(boardRole.Users[:i], boardRole.Users[i+1:]...)
		}
	}
	if query := r.storage.Save(&boardRole); query.Error != nil {
		return models.BoardRole{}, query.Error
	}
	return boardRole, nil
}

func (r *boardRoleRepo) CheckIfHaveAccess(boardId string, userId string) (bool, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("board_id = ?", boardId).First(&boardRole); query.Error != nil {
		return false, query.Error
	}
	for _, user := range boardRole.Users {
		if user.ID == userId {
			return true, nil
		}
	}
	return false, nil
}

func (r *boardRoleRepo) CheckIfHaveAccessToBoardRole(boardRoleId string, userId string) (bool, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", boardRoleId).First(&boardRole); query.Error != nil {
		return false, query.Error
	}
	for _, user := range boardRole.Users {
		if user.ID == userId {
			return true, nil
		}
	}
	return false, nil
}

func (r *boardRoleRepo) GetUsers(boardRoleId string) ([]models.User, error) {
	boardRole := models.BoardRole{}
	if query := r.storage.Where("id = ?", boardRoleId).First(&boardRole); query.Error != nil {
		return []models.User{}, query.Error
	}
	return boardRole.Users, nil
}