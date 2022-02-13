package server

import (
	"e_learning_sample/app/controller"
	"e_learning_sample/app/session"
	"e_learning_sample/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// 設定ファイルの読み込み
	conf := config.GetConfig()
	rootPath := conf.GetString("root_path")

	// 静的ファイルの読み込み
	router.LoadHTMLGlob(rootPath + "/templates/*.html")
	router.Static("/css", rootPath+"/static/css")
	router.Static("/js", rootPath+"/static/js")

	// セッションの作成
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	{
		router.GET("/", controller.IndexAction)
		router.GET("/login", controller.GetLoginAction)
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
