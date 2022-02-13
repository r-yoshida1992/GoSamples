package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexAction(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/logout")
}
