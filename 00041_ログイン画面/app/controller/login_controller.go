package controller

import (
	"00041_login/app/service"
	"00041_login/app/session"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexAction(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func PostLoginAction(c *gin.Context) {
	errMsg := service.LoginCheck(c)
	if len(errMsg) == 0 {
		session.Login(c)
		c.Redirect(http.StatusFound, "/home")
	} else {
		fmt.Println(errMsg)
		c.HTML(200, "login.html", gin.H{
			"errors": errMsg,
		})
	}
}
