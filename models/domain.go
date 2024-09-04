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

// AfterCreate 之后也会调用 AfterSave
func (d *Domain) AfterSave(tx *gorm.DB) error {
	if d.Mode == "cf" {
		api, ctx := utils.InitCfApi(d.ApiKey)
		if d.AccountID == "" {
			d.loadAccountID(api, ctx)
			tx.Save(d)
		}
		if d.ListID == "" {
			d.loadListID(api, ctx)
			tx.Save(d)
		}
	} else if d.Mode == "local" {
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
	// AllowDomain
	if _, err := utils.AllowDomain(d.Name); err != nil {
		log.Println(err)
	}
	if _, err := utils.SaveRules(); err != nil {
		log.Println(err)
	}

	tx.Where("domain_id = ?", d.ID).Delete(&Whitelist{})
	return
}

func (d *Domain) loadAccountID(api *cloudflare.API, ctx context.Context) {
	if d.ApiKey == "" || d.Mode == "local" {
		return
	}
	// init AccountID
	params := cloudflare.AccountsListParams{}
	accounts, _, err := api.Accounts(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	// 默认获取第一个账号ID
	d.AccountID = accounts[0].ID
}

func (d *Domain) loadListID(api *cloudflare.API, ctx context.Context) {
	if d.ApiKey == "" || d.Mode == "local" || d.AccountID == "" {
		return
	}
	// init ListID
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
			break
		}
	}
}

type DomainForm struct {
	Name   string `json:"name" binding:"required"`
	Mode   string `json:"mode" binding:"required"`
	ApiKey string `json:"token"`
	Memo   string `json:"memo"`
}
