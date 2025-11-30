package utils

import (
	"github.com/gin-gonic/gin"
)

type Cookies map[string]string

func GetCookies(c *gin.Context) Cookies {
	cookies := Cookies {}
	for _, cookie := range c.Request.Cookies() {
		if cookie == nil { continue }
		cookies[cookie.Name] = cookie.Value
	}
	return cookies
}
