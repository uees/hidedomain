package services

import (
	"github.com/uees/hidedomain/models"
)

func GetAllDomains(domains *[]models.Domain) error {
	result := DB.Find(&domains)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func QueryDomainByName(name string, domain *models.Domain) error {
	result := DB.Where("name = ?", name).First(&domain)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func HasDomain(name string) (bool, error) {
	var domain models.Domain
	result := DB.Where("name = ?", name).First(&domain)
	if result.RowsAffected > 0 {
		return true, nil
	}

	return false, result.Error
}

func CreateDomain(domain interface{}) error {
	result := DB.Model(&models.Domain{}).Create(&domain)
	return result.Error
}

func UpdateDomainByName(name string, data interface{}) error {
	var domain models.Domain

	if result := DB.Where("name = ?", name).First(&domain); result.Error != nil {
		return result.Error
	}

	if result := DB.Model(&domain).Updates(data); result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteDomainByName(name string) error {
	var domain models.Domain

	if result := DB.Where("name = ?", name).First(&domain); result.Error != nil {
		return result.Error
	}

	if result := DB.Delete(&domain); result.Error != nil {
		return result.Error
	}

	return nil
}
