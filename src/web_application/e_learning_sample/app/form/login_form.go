package form

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

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

func LoginFormValidation(form LoginForm, errMessages []string) []string {
	// ユーザー名の検証
	regexp.MustCompile("[a-z][A-Z][0-9]").Match([]byte(form.UserName))

	return errMessages
}
