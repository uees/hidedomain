package services

import (
	"encoding/json"
	"log"
	"net"

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

func GetAllResolvedProxies(proxies *[]models.Proxyitem) error {
	allProxies := []models.Proxyitem{}
	err := GetAllProxies(&allProxies)
	if err != nil {
		return err
	}

	for _, proxyItem := range allProxies {
		content := make(map[string]interface{})
		err = json.Unmarshal([]byte(proxyItem.Content), &content)
		if err != nil {
			log.Panicln("json.Unmarshal content error", proxyItem.Memo)
			continue
		}
		domain := content["add"].(string)
		addr, err := net.ResolveIPAddr("ip6", domain)
		if err != nil {
			log.Panicln("net.ResolveIPAddr error", domain)
			continue
		}
		content["add"] = addr.IP.String()
		newContent, _ := json.Marshal(content)
		proxyItem.Content = string(newContent)
		*proxies = append(*proxies, proxyItem)
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
	content := make(map[string]interface{})
	err := json.Unmarshal([]byte(data.Content), &content)
	if err != nil {
		return err
	}
	result := db.Create(&models.Proxyitem{
		Memo:     content["ps"].(string),
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

	content := make(map[string]interface{})
	err := json.Unmarshal([]byte(data.Content), &content)
	if err != nil {
		return err
	}
	theUpdateData := structs.Map(data)
	theUpdateData["Memo"] = content["ps"].(string)
	if result := db.Model(&proxyItem).Updates(theUpdateData); result.Error != nil {
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
