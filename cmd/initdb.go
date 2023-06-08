package main

import (
	"errors"
	"fmt"

	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/services"
	"gorm.io/gorm"
)

func main() {
	services.DB.AutoMigrate(&models.User{}, &models.Whitelist{}, &models.Option{}, &models.Domain{})

	var user models.User
	result := services.DB.First(&user, "username = ?", "admin")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		services.AddUser("admin", "admin@localhost", "yourpass")
	}

	fmt.Println("success!")
}
