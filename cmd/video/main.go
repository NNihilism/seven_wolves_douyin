package main

import (
	"douyin/cmd/video/dao"
	video "douyin/kitex_gen/video/videoservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/biz-demo/easy_note/pkg/mw"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)
func init(){
	dao.Init()
}
func main() {
	addr,_ := net.ResolveTCPAddr("tcp",consts.BaseIP+consts.VideoServerPort)
	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServerName}),
	)


	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
