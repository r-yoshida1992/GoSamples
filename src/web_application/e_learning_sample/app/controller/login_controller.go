package controller

import (
	"e_learning_sample/app/service"
	"e_learning_sample/app/session"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetLoginAction
// login
func GetLoginAction(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func PostLoginAction(c *gin.Context) {
	errMessages := service.LoginCheck(c)
	if len(errMessages) == 0 {
		session.Login(c)
		c.Redirect(http.StatusFound, "/home")
	} else {
		fmt.Println(errMessages)
		c.HTML(200, "login.html", gin.H{
			"errors": errMessages,
		})
	}
}
