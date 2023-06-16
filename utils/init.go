package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Conf *Config = new(Config)
	DB   *gorm.DB
)

func init() {
	var err error

	err = Conf.LoadConf()
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(sqlite.Open(Conf.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
