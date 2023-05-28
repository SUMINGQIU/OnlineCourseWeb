package middleware

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// uid, _ := c.Cookie("user_id")
		session := sessions.Default(c)
		uid := session.Get("user_id")
		// session中保存所有值，但是cookie可以很多独立，通过session_secret来进行加密，不然用户可以自己修改cookie成为管理员
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.Response{
			Status: 401,
			Msg:    "您好，需要登录",
		})
		c.Abort()
	}
}
