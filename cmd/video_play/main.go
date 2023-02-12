package main

import (
	"douyin/pkg/consts"
	"fmt"
	"os"
)
import "log"
import "net/http"

const  (
	videoUrl = "/douyin/video/"
	imageUrl = "/douyin/image/"
)

type videoHandler struct{}
type imageHandler struct{}
func (vh *videoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//  日期/视频名字
	path := r.URL.Path[len(videoUrl):]
	fmt.Println("视频路径："+path)
	data, err := os.ReadFile(consts.VideoStorePathPrefix + path)
	if err != nil {
		log.Printf("Error with path %s: %v", path, err)
		w.WriteHeader(404)
		w.Write([]byte("404"))
		return
	}
	w.Header().Add("Accept-Ranges", "bytes")
	w.Header().Add("Content-Type", "video/mp4")
	w.Write(data)
}
func (vh *imageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//  日期/视频名字
	path := r.URL.Path[len(imageUrl):]
	fmt.Println("图片路径"+path)
	data, err := os.ReadFile(consts.ImageStorePathPrefix + path)
	if err != nil {
		log.Printf("Error with path %s: %v", path, err)
		w.WriteHeader(404)
		w.Write([]byte("404"))
		return
	}
	w.Header().Add("Content-Type", "image/jpeg")
	w.Write(data)
}
/**
 * @Author:lemon
 * @Desc:播放视频
 */
func main() {
	http.Handle(videoUrl, new(videoHandler))
	http.Handle(imageUrl,new(imageHandler))
	err := http.ListenAndServe(consts.VideoAndImageServicePort, nil)
	if err != nil {
		fmt.Println("视频流服务启动失败")
		return
	}
}
