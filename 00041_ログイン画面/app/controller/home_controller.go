package controller

import (
	"00041_login/app/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHomeAction(c *gin.Context) {
	s := sessions.Default(c)
	c.HTML(200, "home.html", gin.H{
		"name": s.Get("userName"),
	})
}

func LogoutAction(c *gin.Context) {
	session.Logout(c)
	c.Redirect(http.StatusSeeOther, "/login")
}
