package main

import (
	like "douyin/kitex_gen/like/likeservice"
	"log"
)

func main() {
	svr := like.NewServer(new(LikeServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
