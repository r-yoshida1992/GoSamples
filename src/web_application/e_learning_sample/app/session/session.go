package session

import (
	"e_learning_sample/app/form"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login ログイン時のセッション保存処理
func Login(c *gin.Context) {
	loginForm := form.ParseLoginForm(c)
	session := sessions.Default(c)
	session.Set("userName", loginForm.UserName)
	session.Save()
}

// Logout ログアウト時のセッション破棄処理
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// Check 認証が必要なページでのセッションチェック
func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userName := session.Get("userName")
		if userName == nil {
			c.Redirect(http.StatusFound, "/login") // 302でリダイレクト
			c.Abort()
		} else {
			c.Next()
		}
	}
}
