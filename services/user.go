package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/utils"
	"gorm.io/gorm"
)

func AddUser(username, email, password string) error {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if result.RowsAffected > 0 {
		return fmt.Errorf("user %s is already exists", username)
	}

	hash, err := utils.PasswordHash(password)
	if err != nil {
		return err
	}
	user = models.User{
		Username:      username,
		Email:         email,
		Password:      hash,
		EmailVerified: false,
		Role:          "member",
	}
	result = db.Create(&user)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

var Login = func(info *models.LoginInfo) (bool, error) {
	var user models.User
	result := db.Where("username = ?", info.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}

	if !utils.PasswordVerify(info.Password, user.Password) {
		return false, errors.New("the password is incorrect")
	}

	user.LoginAt = time.Now()
	user.LoginIP = info.IP
	user.LoginUA = info.UserAgent
	db.Save(user)

	return true, nil
}

var DeleteUser = func(username string) (bool, error) {
	result := db.Where("username = ?", username).Delete(&models.User{})
	if result.RowsAffected > 0 {
		return true, nil
	}

	return false, result.Error
}
