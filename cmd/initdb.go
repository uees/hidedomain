package main

import (
	"errors"
	"fmt"

	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/services"
	"github.com/uees/hidedomain/utils"
	"gorm.io/gorm"
)

func main() {
	utils.DB.AutoMigrate(&models.User{}, &models.Whitelist{}, &models.Option{}, &models.Domain{}, &models.Proxyitem{})

	var user models.User
	result := utils.DB.First(&user, "username = ?", "admin")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		services.AddUser("admin", "admin@localhost", "yourpass")
	}

	fmt.Println("success!")
}
