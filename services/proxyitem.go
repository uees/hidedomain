package services

import (
	"github.com/fatih/structs"
	"github.com/uees/hidedomain/models"
)

func GetAllProxies(proxies *[]models.Proxyitem) error {
	result := db.Find(proxies)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func QueryProxyItem(id uint, proxyItem *models.Proxyitem) error {
	result := db.Where("id = ?", id).First(proxyItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateProxyItem(data *models.ProxyitemForm) error {
	result := db.Create(&models.Proxyitem{
		Memo:     data.Memo,
		Content:  data.Content,
		Protocol: data.Protocol,
	})
	return result.Error
}

func UpdateProxyItem(id uint, data *models.ProxyitemForm) error {
	var proxyItem models.Proxyitem

	if result := db.Where("id = ?", id).First(&proxyItem); result.Error != nil {
		return result.Error
	}

	if result := db.Model(&proxyItem).Updates(structs.Map(data)); result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteProxyItem(id uint) error {
	var proxyItem models.Proxyitem

	if result := db.Where("id = ?", id).First(&proxyItem); result.Error != nil {
		return result.Error
	}

	if result := db.Delete(&proxyItem); result.Error != nil {
		return result.Error
	}

	return nil
}
