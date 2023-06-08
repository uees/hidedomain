package models

type Option struct {
	ID    uint   `gorm:"primaryKey" json:"-"`
	Name  string `gorm:"type:varchar(256);not null;uniqueIndex" json:"name"`
	Value string `gorm:"type:varchar(256);" json:"value"`
	Memo  string `gorm:"type:varchar(256);" json:"memo"`
}
