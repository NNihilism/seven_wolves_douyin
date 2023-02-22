package main

import (
	"douyin/cmd/follows/dal"
	follows "douyin/kitex_gen/follows/followservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func init() {
	dal.Init()
}

func main() {

	addr, err := net.ResolveTCPAddr(consts.TCP, consts.FollowServicePort)
	if err != nil {
		panic(err)
	}
	svr := follows.NewServer(new(FollowServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.FollowServiceName}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
