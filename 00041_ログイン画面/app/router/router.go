package router

import (
	"00041_login/app/controller"
	"00041_login/app/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// 静的ファイルの読み込み
	router.LoadHTMLGlob("templates/*.html")
	//router.Static("/css", "/static/css")
	//router.Static("/js", "/static/js")

	// セッションの作成
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	{
		router.GET("/", controller.IndexAction)
		router.GET("/login", controller.IndexAction)
		router.POST("/login", controller.PostLoginAction)
	}

	// ログイン後じゃないと入れないものは以下に定義
	router.Use(session.Check())
	{
		router.GET("/logout", controller.LogoutAction)
		router.GET("/home", controller.GetHomeAction)
	}
	return router
}
