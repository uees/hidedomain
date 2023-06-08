package services

import (
	"errors"

	"github.com/uees/hidedomain/models"
)

func GetWhiteListByDomain(domainName string, whitelists *[]models.Whitelist) error {
	var domain models.Domain
	if err := QueryDomainByName(domainName, &domain); err != nil {
		return err
	}

	if domain.Mode == "cf" {
		return errors.New("cf mode")
	}

	if err := DB.Where("domain_id = ?", domain.ID).Find(whitelists).Error; err != nil {
		return err
	}

	return nil
}

func ClearWhiteListByDomain(domainName string) error {
	var domain models.Domain
	if err := QueryDomainByName(domainName, &domain); err != nil {
		return err
	}

	if domain.Mode == "cf" {
		return errors.New("cf mode")
	}

	if err := DB.Delete(&models.Whitelist{}, "domain_id = ?", domain.ID).Error; err != nil {
		return err
	}

	return nil
}

func AddIPRule(domainName string, r *models.RuleForm) error {
	var domain models.Domain
	if err := QueryDomainByName(domainName, &domain); err != nil {
		return err
	}

	if domain.Mode == "cf" {
		return errors.New("cf mode")
	}

	rule := models.Whitelist{DomainID: domain.ID, Ip: r.Ip, Memo: r.Memo}
	if err := DB.Create(&rule).Error; err != nil {
		return err
	}

	return nil
}

func UpdateIPRule(id string, r *models.RuleForm) error {
	result := DB.Model(models.Whitelist{}).Where("id = ?", id).Updates(*r)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteIPRule(id string) error {
	if err := DB.Delete(&models.Whitelist{}, id).Error; err != nil {
		return err
	}

	return nil
}
