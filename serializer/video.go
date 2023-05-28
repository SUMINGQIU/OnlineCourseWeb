package serializer

import "singo/model"

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

//build a video
func BuildVideo(item model.Video) Video {
	user, _ := model.GetUser(item.UserID)
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:       item.URL,
		Avatar:    item.AvatarURL(),
		User:      BuildUser(user),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
