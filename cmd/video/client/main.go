package main

import (
	"bufio"
	"context"
	video2 "douyin/kitex_gen/video"
	video "douyin/kitex_gen/video/videoservice"
	//"douyin/pkg/consts"
	"fmt"
	"github.com/cloudwego/kitex/client"
	//etcd "github.com/kitex-contrib/registry-etcd"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	//r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	//if err != nil {
	//	panic(err)
	//}
	client, err := video.NewClient("video_module", client.WithHostPorts("0.0.0.0:8083"))

	//client, err := video_module.NewClient("video_module", client.WithResolver(r))

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("C:\\Users\\86176\\Desktop\\test.mp4")
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	videoData := make([]byte, 0)
	buf := make([]byte, 1024*1024)

	for {
		len, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
			// 处理异常
			return
		}
		if len == 0 {
			break
		}
		videoData = append(videoData, buf...)
	}

	req := video2.PublishActionRequest{
		Token: "666",
		Data:  videoData,
		Title: "美女视频",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	resp, err := client.PublishVideo(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
