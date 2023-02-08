package service

import (
	"context"
	"douyin/cmd/video/dao"
	"douyin/cmd/video/pack"
	"douyin/kitex_gen/video"
)

type VideoService struct {
	ctx context.Context
}
func NewVideoService(ctx context.Context)*VideoService{
	return &VideoService{ctx: ctx}
}

func (service *VideoService)GetFeed(req *video.FeedRequest)([]*video.Video,error){
	var resVideos []*video.Video
	videosDbs, err := dao.GetFeed(service.ctx, req.GetLatestTime(), req.GetToken())
	if err!=nil{
		return nil,err
	}
	resVideos = pack.CovertList(videosDbs)
	return resVideos,nil
}
func (service *VideoService)PublishVideo(req *video.PublishActionRequest){

}

