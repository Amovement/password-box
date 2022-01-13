package bcrypt_utils

import (
	bcrypt "golang.org/x/crypto/bcrypt"
)

func GeneratePassWd(src string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)
}
