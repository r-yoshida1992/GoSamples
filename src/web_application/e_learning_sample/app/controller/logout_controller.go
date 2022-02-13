package controller

import (
	"e_learning_sample/app/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutAction(c *gin.Context) {
	session.Logout(c)
	c.Redirect(http.StatusSeeOther, "/login")
}
