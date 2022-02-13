package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetHomeAction(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(200, "home.html", gin.H{
		"name": session.Get("userName"),
	})
}
