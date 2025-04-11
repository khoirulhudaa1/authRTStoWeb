package libraries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCookieUser(c *gin.Context) string {
	user, err := c.Cookie("user")
	if err != nil {
		return "Guest"
	}
	return user
}

func SetCookieUser(c *gin.Context, username string) {
	c.SetCookie("user", username, 3600, "/", "127.0.0.1", false, true)
}

func ClearCookieUser(c *gin.Context) {
	c.SetCookie("user", "", -1, "/", "127.0.0.1", false, true)
}

func IsUserLoggedIn(c *gin.Context) bool {
	user, err := c.Cookie("user")
	return err == nil && user != ""
}

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsUserLoggedIn(c) {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
