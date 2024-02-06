package models

// Proxyitem 订阅代理项目表
type Proxyitem struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Content  string `gorm:"type:text;" json:"content"`
	Memo     string `gorm:"type:varchar(256);" json:"memo"`
	Protocol string `gorm:"type:varchar(256);" json:"protocol"`
}

type ProxyitemForm struct {
	Content  string `json:"content" binding:"required"`
	Memo     string `json:"memo" binding:"required"`
	Protocol string `json:"protocol"`
}
