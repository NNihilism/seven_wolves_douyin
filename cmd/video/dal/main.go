package main

import (
	"douyin/cmd/video/dal/db"
	"fmt"
)

// 测试用 没啥用
func main() {
	Init()

	var video db.Video
	db.DB.First(&video)

	fmt.Println(video)
}
