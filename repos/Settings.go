package repos

type settingsRepo struct {
	storage *gorm.DB
}

// migrate default settings
func (r *settingsRepo) MigrateDefaultSettings() error {
	// check if default settings already exist
	var settings []models.Settings
	if err := r.storage.Find(&settings).Error; err != nil {
		return err
	}
	// insert default settings
	if len(settings) == 0 {
		for _, setting := range models.DefaultSettings {
			if err := r.storage.Create(&setting).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

//update setting
func (r *settingsRepo) Update(setting models.Settings) error {
	// prevent updating any 'Admin' settings
	if setting.Name == "Admin" {
		return errors.New("cannot update Admin settings")
	}
	return r.storage.Save(&setting).Error
}

func (r *settingsRepo) FindByName(name string) (models.Settings, error) {
	setting := models.Settings{}
	if err := r.storage.Find(&setting).Where("name =?", name).Error; err != nil {
		return setting, err
	}
	return setting, nil
}

func (r *settingsRepo) FindAll() ([]models.Settings, error) {
	settings := []models.Settings{}
	if err := r.storage.Find(&settings).Error; err != nil {
		return settings, err
	}
	return settings, nil
}

func (r *settingsRepo) Create(setting models.Settings) (models.Settings, error) {
	if err := r.storage.Create(&setting).Error; err != nil {
		return models.Settings{}, err
	}
	return setting, nil
}

func (r *settingsRepo) Delete(id string) error {
	// prevent deleting any 'Admin' settings
	setting := models.Settings{}
	if err := r.storage.Find(&setting).Where("id =?", id).Error; err != nil {
		return err
	}
	if setting.Name == "Admin" {
		return errors.New("cannot delete Admin settings")
	}
	return r.storage.Delete(&setting).Error
}

func newSettingsRepo(storage *gorm.DB) *settingsRepo {
	return &settingsRepo{
		storage: storage,
	}
}