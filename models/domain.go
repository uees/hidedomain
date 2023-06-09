package models

import (
	"github.com/uees/hidedomain/utils"
	"gorm.io/gorm"
)

type Domain struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(256);not null;uniqueIndex" json:"name"`
	Mode      string `gorm:"type:varchar(32);" json:"mode"` // local or cf
	ZoneID    string `gorm:"type:varchar(64);" json:"-"`    // do not save on server
	AccountID string `gorm:"type:varchar(64);" json:"-"`    // do not save on server
	ApiKey    string `gorm:"type:varchar(64);" json:"-"`    // do not save on server
	Memo      string `gorm:"type:varchar(256);" json:"memo"`
}

func (d *Domain) AfterCreate(tx *gorm.DB) (err error) {
	// DenyDomain
	if d.Mode == "local" {
		utils.DenyDomain(d.Name)
	}
	return
}

func (d *Domain) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Where("domain_id = ?", d.ID).Delete(&Whitelist{})
	return
}

func (d *Domain) AfterDelete(tx *gorm.DB) (err error) {
	// AllowDomain
	utils.AllowDomain(d.Name)
	return
}

type DomainForm struct {
	Name string `json:"name" binding:"required"`
	Mode string `json:"mode" binding:"required"`
	Memo string `json:"memo"`
}
