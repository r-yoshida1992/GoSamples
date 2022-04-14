package service

import (
	"00041_login/app/form"
	"github.com/gin-gonic/gin"
)

func LoginCheck(c *gin.Context) []string {

	// ログイン情報を取得
	loginForm := form.ParseLoginForm(c)

	var errMsg []string

	if loginForm.UserName == "" {
		errMsg = append(errMsg, "ユーザ名が入力されていません")
	}

	if loginForm.Password == "" {
		errMsg = append(errMsg, "パスワードが入力されていません")
	}

	return errMsg
}
