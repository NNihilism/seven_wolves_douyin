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
	resp,err = videoClient.GetFeed(ctx,req)
	if err != nil {
		log.Fatal(err)
		return &video.FeedResponse{StatusCode: -1}, err
	}
	return resp, nil
}

func PublishVideo(ctx context.Context, req *video.PublishActionRequest) (r *video.PublishActionResponse, err error) {
	resp, err := videoClient.PublishVideo(ctx, req)
	if err != nil {
		log.Fatal(err)
		return &video.PublishActionResponse{StatusCode: -1}, err
	}
	return resp, nil
}
