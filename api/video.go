package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo

func CreateVideo(c *gin.Context) {
	// handle the unexpected things
	user := CurrentUser(c)
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// Show Video

func ShowVideo(c *gin.Context) {
	// handle the unexpected things
	service := service.ShowVideoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)

}

// List Video

func ListVideo(c *gin.Context) {
	service := service.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// update the video

func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// delete the video
func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)

}
