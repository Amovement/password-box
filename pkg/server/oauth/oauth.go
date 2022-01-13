package oauth

import (
	"github.com/Amovement/password-box/pkg/dao"
	bcryt_utils "github.com/Amovement/password-box/pkg/utils/bcrypt_utils"
)

func CheckUserNameAndPassword(username string, password string) bool {
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		return false
	}
	input_pwd, err := bcryt_utils.GeneratePassWd(password)
	if err != nil {
		return false
	}
	user_pwd, err := bcryt_utils.GeneratePassWd(user.Password)
	if err != nil {
		return false
	}

	if string(input_pwd) != string(user_pwd) {
		return false
	}
	return true
}
