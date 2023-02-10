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

	var err error
	videoClient, err = videoservice.NewClient("video_module", client.WithHostPorts(consts.VideoServerHost+consts.VideoServerPort))

	if err != nil {
		log.Fatal(err)
		return
	}
}
func PublishVideo(ctx context.Context, req *video.PublishActionRequest) (r *video.PublishActionResponse, err error) {
	resp, err := videoClient.PublishVideo(ctx, req)

	if err != nil {
		log.Fatal(err)
		return &video.PublishActionResponse{StatusCode: -1}, err
	}

	return resp, nil
}
