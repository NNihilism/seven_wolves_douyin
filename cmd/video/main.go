package main

import (
	"douyin/cmd/video/dao"
	video "douyin/kitex_gen/video/videoservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)
func init(){
	dao.Init()
}
func main() {
	addr,_ := net.ResolveTCPAddr("tcp",consts.BaseIP+consts.VideoServicePort)
	svr := video.NewServer(new(VideoServiceImpl),server.WithServiceAddr(addr))

	// init db
	dao.Init()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}