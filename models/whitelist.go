package models

import (
	"log"

	"github.com/uees/hidedomain/utils"
)

// Whitelist IP 白名单，仅仅保存 local 模式的数据
type Whitelist struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	DomainID uint   `json:"domain_id"`
	Domain   Domain `gorm:"constraint:OnDelete:SET NULL;" json:"domain"`
	Ip       string `gorm:"type:varchar(256);" json:"ip"`
	Memo     string `gorm:"type:varchar(256);" json:"memo"`
}

func (wl *Whitelist) AllowIP(domain string) (err error) {
	// AllowIP
	if _, err := utils.AllowIP(domain, wl.Ip); err != nil {
		log.Println(err)
	}
	if _, err := utils.SaveRules(); err != nil {
		log.Println(err)
	}
	return
}

func (wl *Whitelist) RemoveIP(domain string) (err error) {
	// RemoveIP
	if _, err := utils.RemoveIP(domain, wl.Ip); err != nil {
		log.Println(err)
	}
	if _, err := utils.SaveRules(); err != nil {
		log.Println(err)
	}
	return
}

type RuleForm struct {
	Ip   string
	Memo string
}

// 更新时只能更新备注信息
type RuleUpdateForm struct {
	Memo string
}

// 统一 local data 和 cloudflare data
type Rule struct {
	Key    string `json:"key"`
	Domain string `json:"domain"`
	Ip     string `json:"ip"`
	Memo   string `json:"memo"`
}
