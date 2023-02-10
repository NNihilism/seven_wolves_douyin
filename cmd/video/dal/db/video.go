package db

import (
	"context"
	"douyin/pkg/consts"
	"time"
)

type Video struct {
	Id          int64     `json:"id"`
	AnchorId    int64     `json:"anchor_id"`
	PlayUrl     string    `json:"play_url"`
	CoverUrl    string    `json:"cover_url"`
	Title       string    `json:"title"`
	PublishTime time.Time `json:"publish_time"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

func QueryVideoById(ctx context.Context, videoId int64) (video Video, err error) {

	return Video{}, nil
}

func CreateVideo(video *Video) error {
	result := DB.Create(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
