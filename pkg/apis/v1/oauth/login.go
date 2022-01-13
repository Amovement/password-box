package oauth

import (
	"github.com/Amovement/password-box/pkg/models"
	OauthServer "github.com/Amovement/password-box/pkg/server/oauth"
	"github.com/Amovement/password-box/pkg/utils/result"
	"github.com/gin-gonic/gin"
)

type OauthRegister struct {
}

func (s *OauthRegister) Login(c *gin.Context) {
	result := result.NewResult(c)

	var user *models.User
	err := c.BindJSON(&user)
	if err != nil {
		result.Error(500, "Data Loss")
		return
	}
	if OauthServer.CheckUserNameAndPassword(user.Username, user.Password) {

	} else {
		result.Error(500, "username or password not found")
		return
	}

}
