package model

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title  string
	Info   string
	URL    string
	Avatar string
	UserID uint
}

// AvatarURL 封面地址

func (video *Video) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}
