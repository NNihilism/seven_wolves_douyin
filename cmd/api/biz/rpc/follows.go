package rpc

import (
	"context"
	"douyin/kitex_gen/follows"
	"douyin/kitex_gen/follows/followservice"
	"douyin/pkg/consts"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"log"
)

var followClient followservice.Client

func ChangeRelationAction(ctx context.Context, req *follows.RelationActionRequest) (resp *follows.RelationActionResponse, err error) {
	// fmt.Println(followClient)
	resp = new(follows.RelationActionResponse)
	resp, err = followClient.RelationAction(ctx, req)
	// fmt.Println("--------------------------------")
	// fmt.Println(resp)
	code := resp.BaseResp.GetStatusCode()
	if err != nil {
		resp.BaseResp.StatusCode = -1
		resp.BaseResp.StatusMessage = "关注/取关操作失败"
		log.Fatal(err)
		return
	}
	if code == 0 {
		resp.BaseResp.StatusMessage = "关注/取关操作成功"
	} else {
		resp.BaseResp.StatusMessage = "关注/取关操作失败"
	}
	return resp, err
}
func InitClient() {
	c, err := followservice.NewClient(
		consts.FollowServiceName,
		client.WithHostPorts(consts.FollowserverHost+consts.FollowServicePort),
	)
	fmt.Println(consts.FollowserverHost + consts.FollowServicePort)
	if err != nil {
		log.Fatal(err)
	}
	followClient = c
}
