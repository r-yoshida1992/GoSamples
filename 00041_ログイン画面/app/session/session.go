package session

import (
	"00041_login/app/form"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login ログイン時のセッション保存処理
func Login(c *gin.Context) {
	f := form.ParseLoginForm(c)
	s := sessions.Default(c)
	s.Set("userName", f.UserName)
	err := s.Save()
	if err != nil {
		return
	}
}

// Logout ログアウト時のセッション破棄処理
func Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	err := s.Save()
	if err != nil {
		return
	}
}

// Check 認証が必要なページでのセッションチェック
func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		userName := s.Get("userName")
		if userName == nil {
			c.Redirect(http.StatusFound, "/login") // 302でリダイレクト
			c.Abort()
		} else {
			c.Next()
		}
	}
}
