package service

import (
	"errors"
	"singo/model"
	"singo/serializer"
	"strconv"
)

// Show video service: login manage
type ShowVideoService struct {
}

// Show videos
func (service *ShowVideoService) Show(id string) serializer.Response {
	// to prevent the sql injection
	if _, err := strconv.Atoi(id); err != nil {
		return serializer.ParamErr("", errors.New("id不合法"))
	}
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
