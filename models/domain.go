package models

type Domain struct {
	ID        uint   `gorm:"primaryKey" json:"-"`
	Name      string `gorm:"type:varchar(256);not null;uniqueIndex" json:"name"`
	Mode      string `gorm:"type:varchar(32);" json:"mode"` // local or cf
	ZoneID    string `gorm:"type:varchar(64);" json:"-"`
	AccountID string `gorm:"type:varchar(64);" json:"-"`
	ApiKey    string `gorm:"type:varchar(64);" json:"-"`
	Memo      string `gorm:"type:varchar(256);" json:"memo"`
}

type DomainForm struct {
	Name      string `json:"domain" binding:"required"`
	Mode      string `json:"mode" binding:"required"`
	ZoneID    string `json:"zone_id"`
	AccountID string `json:"account_id"`
	ApiKey    string `json:"api_key"`
	Memo      string `json:"memo"`
}
