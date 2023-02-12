package main

import (
	"context"
	"douyin/cmd/video/pack"
	"douyin/cmd/video/service"
	video "douyin/kitex_gen/video"
	"log"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	resp = new(video.FeedResponse)
	videoList, err := service.NewVideoService(ctx).GetFeed(req)
	resp.SetVideoList(pack.CovertList(videoList))
	resp.SetStatusCode(0)
	nextTime := videoList[len(videoList)-1].CreatedAt.Unix()
	resp.SetNextTime(&nextTime)
	return resp, nil
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(video.PublishActionResponse)
	msg := "ok"
	resp.SetStatusCode(0)
	resp.SetStatusMsg(&msg)
	err = service.NewVideoService(ctx).PublishVideo(req)
	if err != nil {
		msg = "publish error"
		resp.SetStatusCode(-1)
		log.Println("发布失败")
		return resp, err
	}
	return resp, nil
}

// GetPublishVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishVideoList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
