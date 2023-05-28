package service

import (
	"singo/model"
	"singo/serializer"
)

// create video service: login manage
type CreateVideoService struct {
	Title  string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info   string `form:"info" json:"info" binding:"max=3000"`
	URL    string `form:"url" json:"url"`
	Avatar string `form:"avatar" json:"avatar"`
}

// Create videos
func (service *CreateVideoService) Create(user *model.User) serializer.Response {
	video := model.Video{
		Title:  service.Title,
		Info:   service.Info,
		URL:    service.URL,
		Avatar: service.Avatar,
		UserID: user.ID,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "账号或密码错误",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
