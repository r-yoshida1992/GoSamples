package form

import "github.com/gin-gonic/gin"

type LoginForm struct {
	UserName string
	Password string
}

func ParseLoginForm(c *gin.Context) LoginForm {
	var form LoginForm
	form.UserName = c.PostForm("userId")
	form.Password = c.PostForm("password")
	return form
}
