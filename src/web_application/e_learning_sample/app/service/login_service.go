package service

import (
	"e_learning_sample/app/form"
	"github.com/gin-gonic/gin"
)

func LoginCheck(c *gin.Context) []string {

	// ログイン情報を取得
	loginForm := form.ParseLoginForm(c)

	var errMessages []string

	if loginForm.UserName == "" {
		errMessages = append(errMessages, "ユーザ名が入力されていません")
	}

	if loginForm.Password == "" {
		errMessages = append(errMessages, "パスワードが入力されていません")
	}

	return errMessages
}
