package api

import (
	"singo/serializer"
	"github.com/gin-gonic/gin"
)
// CreateVideo 
func CreateVideo(c *gin.Context) { 
	// c is defined by gin.context which means we dont need to give any arg
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "创建成功",
	})
}