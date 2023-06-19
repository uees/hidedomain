package models

import (
	"context"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/uees/hidedomain/utils"
	"gorm.io/gorm"
)

type Domain struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(256);not null;uniqueIndex" json:"name"`
	Mode      string `gorm:"type:varchar(32);" json:"mode"`  // local or cf
	ListID    string `gorm:"type:varchar(64);" json:"-"`     // list_id
	AccountID string `gorm:"type:varchar(64);" json:"-"`     // account_id
	ApiKey    string `gorm:"type:varchar(64);" json:"token"` // cf token
	Memo      string `gorm:"type:varchar(256);" json:"memo"`
}

func (d *Domain) BeforeSave(tx *gorm.DB) error {
	if d.ApiKey != "" {
		api, ctx := utils.InitCfApi(d.ApiKey)
		d.loadAccountID(tx, api, ctx)
		d.loadListID(tx, api, ctx)
	}
	return nil
}

func (d *Domain) AfterSave(tx *gorm.DB) error {
	if d.Mode == "local" {
		if _, err := utils.AllowDomain(d.Name); err != nil {
			log.Println(err)
		}
		if _, err := utils.DenyDomain(d.Name); err != nil {
			log.Println(err)
		}
		if _, err := utils.SaveRules(); err != nil {
			log.Println(err)
		}
	}

	return nil
}

func (d *Domain) BeforeCreate(tx *gorm.DB) error {
	return d.BeforeSave(tx)
}

func (d *Domain) AfterCreate(tx *gorm.DB) (err error) {
	// DenyDomain
	if d.Mode == "local" {
		if _, err := utils.DenyDomain(d.Name); err != nil {
			log.Println(err)
		}
		if _, err := utils.SaveRules(); err != nil {
			log.Println(err)
		}
	}
	return
}

func (d *Domain) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Where("domain_id = ?", d.ID).Delete(&Whitelist{})
	return
}

func (d *Domain) AfterDelete(tx *gorm.DB) (err error) {
	// AllowDomain
	if _, err := utils.AllowDomain(d.Name); err != nil {
		log.Println(err)
	}
	if _, err := utils.SaveRules(); err != nil {
		log.Println(err)
	}
	return
}

func (d *Domain) loadAccountID(tx *gorm.DB, api *cloudflare.API, ctx context.Context) {
	if d.ApiKey == "" || d.Mode == "local" {
		return
	}
	// init AccountID
	if d.AccountID == "" {
		params := cloudflare.AccountsListParams{}
		accounts, _, err := api.Accounts(ctx, params)
		if err != nil {
			log.Fatal(err)
		}

		// 默认获取第一个账号ID
		d.AccountID = accounts[0].ID
		tx.Save(d)
	}
}

func (d *Domain) loadListID(tx *gorm.DB, api *cloudflare.API, ctx context.Context) {
	if d.ApiKey == "" || d.Mode == "local" || d.AccountID == "" {
		return
	}
	// init ListID
	if d.ListID == "" {
		var rc = &cloudflare.ResourceContainer{
			Level:      cloudflare.AccountRouteLevel,
			Identifier: d.AccountID,
		}
		var params = cloudflare.ListListsParams{}
		lists, err := api.ListLists(ctx, rc, params)
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			// 默认获取 "my_ip_list"
			if list.Name == "my_ip_list" {
				d.ListID = list.ID
				tx.Save(d)
				break
			}
		}
	}
}

type DomainForm struct {
	Name   string `json:"name" binding:"required"`
	Mode   string `json:"mode" binding:"required"`
	ApiKey string `json:"token"`
	Memo   string `json:"memo"`
}
