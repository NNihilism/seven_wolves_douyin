package main

import (
	follows "douyin/kitex_gen/follows/followservice"
	"log"
)

func main() {
	svr := follows.NewServer(new(FollowServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
