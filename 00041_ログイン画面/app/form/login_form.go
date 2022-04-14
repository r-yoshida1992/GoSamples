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
	return LoginForm{c.PostForm("userId"), c.PostForm("password")}
}

func LoginFormValidation(form LoginForm, errMsg []string) []string {
	// ユーザー名の検証
	regexp.MustCompile("[a-z][A-Z][0-9]").Match([]byte(form.UserName))

	return errMsg
}
