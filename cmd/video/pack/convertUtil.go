package pack

import (
	"douyin/cmd/video/dao"
	"douyin/kitex_gen/video"
)

func ConvertOne(v *dao.Video) *video.Video {
	return &video.Video{
		Id:       int64(v.Id),
		PlayUrl:  v.PlayUrl,
		CoverUrl: v.CoverUrl,
		Title:    v.Title,
	}
}
func CovertList(vs []*dao.Video) []*video.Video {
	var videos []*video.Video
	for _, item := range vs {
		if tmp := ConvertOne(item); tmp != nil {
			videos = append(videos, tmp)
		}
	}
	return videos
}
