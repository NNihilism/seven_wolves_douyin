package dao

import (
	"context"
	"time"
)

type Video struct {
	Id        int64     `json:"id"`
	AuthorId  int64     `json:"author_id"`
	PlayUrl   string    `json:"play_url"`
	CoverUrl  string    `json:"cover_url"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (v *Video) TableName() string {
	return "video"
}
func GetFeed(ctx context.Context, latestTime int64) ([]*Video, error) {
	var videos []*Video
	wrapper := make(map[string]interface{})
	if latestTime != 0 {
		wrapper["created_at"] = latestTime
	}
	err := DB.WithContext(ctx).Where(wrapper).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
func CreateVideo(ctx context.Context, v *Video) error {
	err := DB.WithContext(ctx).Create(v).Error
	return err
}
