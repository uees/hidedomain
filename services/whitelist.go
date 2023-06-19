package services

import (
	"context"
	"log"
	"strconv"

	"github.com/cloudflare/cloudflare-go"
	"github.com/fatih/structs"
	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/utils"
)

var (
	api *cloudflare.API
	ctx context.Context
)

func GetWhiteListByDomain(domainName string, whitelists *[]models.Rule) error {
	var domain models.Domain
	if err := QueryDomainByName(domainName, &domain); err != nil {
		return err
	}

	if domain.Mode == "cf" {
		api, ctx = utils.InitCfApi(domain.ApiKey)
		rc := &cloudflare.ResourceContainer{
			Level:      cloudflare.AccountRouteLevel,
			Identifier: domain.AccountID,
		}
		params := cloudflare.ListListItemsParams{
			ID: domain.ListID,
		}
		list, err := api.ListListItems(ctx, rc, params)
		if err != nil {
			// 只记录错误，不抛出
			log.Println(err)
			return nil
		}

		// out whitelists
		for _, rule := range list {
			*whitelists = append(*whitelists, models.Rule{
				Key:    rule.ID,
				Domain: domain.Name,
				Ip:     *rule.IP,
				Memo:   rule.Comment,
			})
		}

		return nil
	}

	localData := []models.Whitelist{}
	if err := db.Where("domain_id = ?", domain.ID).Find(&localData).Error; err != nil {
		return err
	}

	// out whitelists
	for _, rule := range localData {
		*whitelists = append(*whitelists, models.Rule{
			Key:    strconv.FormatInt(int64(rule.ID), 10),
			Domain: domainName,
			Ip:     rule.Ip,
			Memo:   rule.Memo,
		})
	}

	return nil
}

func ClearWhiteListByDomain(domainName string) error {
	var domain models.Domain
	if err := QueryDomainByName(domainName, &domain); err != nil {
		return err
	}

	if domain.Mode == "cf" {
		// nothing to do
		log.Println("cf mode not to do clear all list")
		return nil
	}

	// Remove ips
	list := []models.Whitelist{}
	if err := db.Where("domain_id = ?", domain.ID).Find(&list).Error; err != nil {
		return err
	}
	for _, rule := range list {
		rule.RemoveIP(domainName)
	}

	if err := db.Delete(&models.Whitelist{}, "domain_id = ?", domain.ID).Error; err != nil {
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
		api, ctx = utils.InitCfApi(domain.ApiKey)
		rc := &cloudflare.ResourceContainer{
			Level:      cloudflare.AccountRouteLevel,
			Identifier: domain.AccountID,
		}
		params := cloudflare.ListCreateItemParams{
			ID: domain.ListID,
			Item: cloudflare.ListItemCreateRequest{
				IP:      &r.Ip,
				Comment: r.Memo,
			},
		}
		_, err := api.CreateListItem(ctx, rc, params)
		if err != nil {
			// 只记录错误，不抛出
			log.Println(err)
			return nil
		}
		return nil
	}

	// 防止重复
	find := models.Whitelist{}
	result := db.Where("ip = ? AND domain_id = ?", r.Ip, domain.ID).Find(&find)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		log.Printf("Rule %s aready exist\n", find.Ip)
		return nil
	}

	rule := models.Whitelist{DomainID: domain.ID, Ip: r.Ip, Memo: r.Memo}
	if err := db.Create(&rule).Error; err != nil {
		return err
	}

	rule.AllowIP(domainName)

	return nil
}

func UpdateIPRule(domainName string, id string, r *models.RuleUpdateForm) error {
	var domain models.Domain
	if err := QueryDomainByName(domainName, &domain); err != nil {
		return err
	}

	if domain.Mode == "local" {
		result := db.Model(models.Whitelist{}).Where("id = ? AND domain_id = ?", id, domain.ID).Updates(structs.Map(r))
		if result.Error != nil {
			return result.Error
		}
		return nil
	}

	// todo cf
	log.Println("cf mode not to do ip rule update action")
	return nil
}

func DeleteIPRule(domainName string, id string) error {
	var domain models.Domain
	if err := QueryDomainByName(domainName, &domain); err != nil {
		return err
	}

	if domain.Mode == "local" {
		rule := models.Whitelist{}
		db.First(&rule, id)
		if err := rule.RemoveIP(domainName); err != nil {
			return err
		}

		if err := db.Where("id = ? AND domain_id = ?", id, domain.ID).Delete(&models.Whitelist{}).Error; err != nil {
			return err
		}

		return nil
	}

	api, ctx = utils.InitCfApi(domain.ApiKey)
	rc := &cloudflare.ResourceContainer{
		Level:      cloudflare.AccountRouteLevel,
		Identifier: domain.AccountID,
	}
	params := cloudflare.ListDeleteItemsParams{
		ID: domain.ListID,
		Items: cloudflare.ListItemDeleteRequest{
			Items: []cloudflare.ListItemDeleteItemRequest{
				{
					ID: id,
				},
			},
		},
	}
	_, err := api.DeleteListItems(ctx, rc, params)
	if err != nil {
		// 只记录错误，不抛出
		log.Println(err)
		return nil
	}

	return nil
}

func GetIpRule(domainName string, id string, r *models.Rule) error {
	var domain models.Domain
	if err := QueryDomainByName(domainName, &domain); err != nil {
		return err
	}

	if domain.Mode == "local" {
		rule := models.Whitelist{}
		result := db.Where("id = ?", id).Find(&rule)
		if result.Error != nil {
			return result.Error
		}

		// out to r
		*r = models.Rule{
			Key:    strconv.FormatInt(int64(rule.ID), 10),
			Domain: domainName,
			Ip:     rule.Ip,
			Memo:   rule.Memo,
		}
		return nil
	}

	// mode == "cf"
	api, ctx = utils.InitCfApi(domain.ApiKey)
	rc := &cloudflare.ResourceContainer{
		Level:      cloudflare.AccountRouteLevel,
		Identifier: domain.AccountID,
	}
	rule, err := api.GetListItem(ctx, rc, domain.ListID, id)
	if err != nil {
		// 只记录错误，不抛出
		log.Println(err)
		return nil
	}

	// out to r
	*r = models.Rule{
		Key:    rule.ID,
		Domain: domainName,
		Ip:     *rule.IP,
		Memo:   rule.Comment,
	}
	return nil
}
