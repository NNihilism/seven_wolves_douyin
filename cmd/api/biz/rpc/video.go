package rpc

import (
	"context"
	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"log"
	"sync"
)

var (
	videoClient videoservice.Client
	lock        sync.Mutex
)

// 防止服务尚未启动时建立连接导致错误，调用服务时再调用该方法
func GetVideoClient() videoservice.Client {
	// 单例模式创建客户端
	if videoClient == nil {
		if lock.TryLock() {
			if videoClient == nil {
				// todo 后续可优化为从注册中心获取服务信息
				c, err := videoservice.NewClient(
					consts.VideoServerName,
					client.WithHostPorts(consts.VideoServerHost+consts.VideoServerPort),
				)

				if err != nil {
					log.Fatal(err)
				}
				videoClient = c
			}
			lock.Unlock()
		}
	}
	return videoClient
}
func GetFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	vc := GetVideoClient()
	resp, err = vc.GetFeed(ctx, req)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}

func PublishVideo(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	vc := GetVideoClient()
	resp, err = vc.PublishVideo(ctx, req)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}

func GetPublishVideoList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	vc := GetVideoClient()
	resp, err = vc.GetPublishVideoList(ctx, req)
	//if err != nil {
	//	log.Fatal(err)
	//	return resp, err
	//}
	return resp, nil
}
