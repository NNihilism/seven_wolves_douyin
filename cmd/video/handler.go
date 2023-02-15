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
	resp = new(video.PublishActionResponse)
	err = service.NewVideoService(ctx).PublishVideo(req)
	if err != nil {
		log.Fatal("发布失败")
		return &video.PublishActionResponse{StatusCode: -1}, err
	}
	resp.SetStatusCode(0)
	return resp, nil
}

// GetPublishVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishVideoList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...、
	resp = new(video.PublishListResponse)
	videoList, err := service.NewVideoService(ctx).GetPublishVideoList(req)
	if err != nil {
		log.Fatal("获取用户的视频发布列表失败")
		resp.SetStatusCode(-1)
		msg := "获取失败"
		resp.SetStatusMsg(&msg)
		return resp, err
	}
	resp.SetStatusCode(0)
	resp.SetVideoList(pack.CovertList(videoList))

	return resp, nil
}
