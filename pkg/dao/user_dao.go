package dao

import (
	"github.com/Amovement/password-box/pkg/models"
	bcryt_utils "github.com/Amovement/password-box/pkg/utils/bcrypt_utils"
)

func GetUserByUsername(username string) (*models.User, error) {
	db := models.GetDB()
	var user *models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func CreateUser(username string, pwd string) (*models.User, error) {
	db := models.GetDB()
	encode_pwd, _ := bcryt_utils.GeneratePassWd(pwd)
	user := &models.User{
		Username: username,
		Password: string(encode_pwd),
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
