package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Amovement/password-box/pkg/config"
	"github.com/Amovement/password-box/pkg/router/middlewares"
	"github.com/Amovement/password-box/pkg/router/register"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func initMiddleware(r *gin.Engine) {
	cfg := config.GetConfig()
	store := cookie.NewStore([]byte(cfg.App.Session_key))
	log := middlewares.GetLog()
	//跨域中间件
	r.Use(sessions.Sessions("pwdbox_session", store), middlewares.Logger(log), middlewares.Cors, gin.Recovery())
}

func SetupRouterAndGetServer() *http.Server {
	cfg := config.GetConfig()
	r := gin.Default()
	initMiddleware(r)

	register.InitOauthRouter(r)

	GinServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return GinServer
}
