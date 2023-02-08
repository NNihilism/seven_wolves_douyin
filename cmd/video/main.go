package main

import (
	"douyin/cmd/video/dal/db"
	video "douyin/kitex_gen/video/videoservice"
	"log"
)

func main() {
	svr := video.NewServer(new(VideoServiceImpl))

	// init db
	db.Init()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
