package models

// Whitelist IP 白名单，仅仅保存 local 模式的数据
type Whitelist struct {
	ID       uint   `gorm:"primaryKey"`
	DomainID uint   `json:"-"`
	Domain   Domain `gorm:"constraint:OnDelete:SET NULL;" json:"domain"`
	Ip       string `gorm:"type:varchar(256);" json:"ip"`
	Memo     string `gorm:"type:varchar(256);" json:"memo"`
}

type RuleForm struct {
	Ip   string
	Memo string
}
