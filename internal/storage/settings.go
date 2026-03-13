package storage

// GetSettings returns the current settings
func GetSettings() (*Settings, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var settings Settings
	db.Read(func(data *StorageData) {
		settings = data.Settings
	})

	return &settings, nil
}

// UpdateSettings updates the settings
func UpdateSettings(updates map[string]string) (*Settings, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var settings Settings
	err = db.Update(func(data *StorageData) {
		if listLayout, ok := updates["listLayout"]; ok {
			data.Settings.ListLayout = listLayout
		}
		if dateFormat, ok := updates["dateFormat"]; ok {
			data.Settings.DateFormat = dateFormat
		}
		settings = data.Settings
	})

	if err != nil {
		return nil, err
	}

	return &settings, nil
}

// ResetSettings resets settings to defaults
func ResetSettings() (*Settings, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	defaults := DefaultSettings()
	err = db.Update(func(data *StorageData) {
		data.Settings = defaults
	})

	if err != nil {
		return nil, err
	}

	return &defaults, nil
}
