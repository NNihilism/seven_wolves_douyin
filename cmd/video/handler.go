package main

import (
	"context"
	video "douyin/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPublishVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishVideoList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
