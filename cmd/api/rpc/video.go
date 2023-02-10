package rpc

import (
	"context"
	httpVideo "douyin/cmd/api/biz/model/video"
	rpcVideo "douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"log"
)

func InitVideoClient() (videoservice.Client, error) {

	var err error
	videoClient, err := videoservice.NewClient("video_module", client.WithHostPorts(consts.VideoServerHost+consts.VideoServerPort))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return videoClient, nil
}
func PublishVideo(client videoservice.Client, ctx context.Context, httpReq *httpVideo.PublishActionRequest) (r *rpcVideo.PublishActionResponse, err error) {

	// todo 根据httpReq中的token进行鉴权，可能不需要在服务内部写

	// 封装rpc请求对象
	rpcReq := rpcVideo.PublishActionRequest{
		Token: httpReq.Token,
		Data:  httpReq.Data,
		Title: httpReq.Title,
	}
	// 发起rpc调用
	resp, err := client.PublishVideo(ctx, &rpcReq)

	if err != nil {
		log.Fatal(err)
		return &rpcVideo.PublishActionResponse{StatusCode: -1}, err
	}

	return resp, nil
}
