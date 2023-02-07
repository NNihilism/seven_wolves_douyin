package db

import (
	"context"
	"douyin/pkg/consts"
	"time"
)

type Video struct {
	Id          int64
	AnchorId    int64
	PlayUrl     string
	CoverUrl    string
	Title       string
	PublishTime time.Time
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

func QueryVideoById(ctx context.Context, videoId int64) (video Video, err error) {

	return Video{}, nil
}

func CreateVideo(video Video) error {

	return nil
}
