package serializer

import (
	"log"
	"singo/model"
)

// Video 用户序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json: "url"`
	Avatar    string `json: "avatar"`
	User      User   `json:"user"`
	CreatedAt int64  `json:"created_at"`
}

// build a video
func BuildVideo(item model.Video) Video {
	user, _ := model.GetUser(item.UserID)
	avatarUrl, err := item.AvatarURL()
	if err != nil {
		log.Printf("Error getting avatar URL: %v", err)
		// handle the error, for example by setting avatarUrl to a default value
		avatarUrl = ""
	}
	return Video{
		ID:    item.ID,
		Title: item.Title,
		Info:  item.Info,
		URL:   item.URL,
		// Avatar:    item.AvatarURL(),
		Avatar:    avatarUrl,
		User:      BuildUser(user),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		if &item != nil {
			video := BuildVideo(item)
			videos = append(videos, video)
		}

	}
	return videos
}
