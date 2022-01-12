package register

import (
	oauth "github.com/Amovement/password-box/pkg/apis/v1/oauth"
	"github.com/gin-gonic/gin"
)

func InitOauthRouter(r *gin.Engine) {

	oauthController := oauth.OauthRegister{}

	r.POST("/login", oauthController.Login)
}
