package services

import (
	"errors"

	"github.com/uees/hidedomain/models"
	"gorm.io/gorm"
)

func GetOption(name string) (string, error) {
	var op models.Option
	result := db.Where("name = ?", name).First(&op)
	if result.Error != nil {
		return "", result.Error
	}

	return op.Value, nil
}

func SetOption(name string, value string) error {
	var op models.Option
	result := db.Where("name = ?", name).First(&op)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result = db.Create(&models.Option{
				Name:  name,
				Value: value,
			})
			return result.Error
		}

		return result.Error
	}

	op.Value = value
	result = db.Save(&op)
	return result.Error
}

func GetAllOptions() map[string]string {
	var options []models.Option
	db.Find(&options)

	result := map[string]string{}
	for _, option := range options {
		result[option.Name] = option.Value
	}

	return result
}
