package db

import (
	"douyin/pkg/consts"
	"time"
)

type Video struct {
	Id            int64
	AnchorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64 // 点赞业务
	CommentCount  int64 // 评论业务
	IsFavorite    bool  // 点赞业务
	Title         string
	PublishTime   time.Time
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}
