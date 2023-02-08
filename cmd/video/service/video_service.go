package service

import (
	"context"
	"douyin/cmd/video/dao"
	"douyin/kitex_gen/video"
)

type VideoService struct {
	ctx context.Context
}
func NewVideoService(ctx context.Context)*VideoService{
	return &VideoService{ctx: ctx}
}

func (service *VideoService)GetFeed(req *video.FeedRequest)([]*dao.Video,error){
	videosDbs, err := dao.GetFeed(service.ctx, req.GetLatestTime(), req.GetToken())
	if err!=nil{
		return nil,err
	}
	return videosDbs,nil
}
func (service *VideoService)PublishVideo(req *video.PublishActionRequest){

}

