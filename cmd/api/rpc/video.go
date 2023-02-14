package rpc

import (
	"context"
	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"log"
)

var videoClient videoservice.Client

func InitVideoClient() {

	c, err := videoservice.NewClient(
		consts.VideoServerName,
		client.WithHostPorts(consts.VideoServerHost+consts.VideoServerPort),
	)

	if err != nil {
		log.Fatal(err)
	}
	videoClient = c
}
func GetFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp, err = videoClient.GetFeed(ctx, req)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}

func PublishVideo(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	resp, err = videoClient.PublishVideo(ctx, req)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}

func GetPublishVideoList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	resp, err = videoClient.GetPublishVideoList(ctx, req)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}
