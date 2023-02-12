package main

import (
	"douyin/pkg/consts"
	"fmt"
	"io/ioutil"
)
import "log"
import "net/http"

type viewHandler struct{}

func (vh *viewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	fmt.Println(path)
	data, err := ioutil.ReadFile("static/video"+path)
	fmt.Println("请求进来了")
	if err != nil {
		log.Printf("Error with path %s: %v", path, err)
		w.WriteHeader(404)
		w.Write([]byte("404"))
	}
	w.Header().Add("Accept-Ranges","bytes")
	w.Header().Add("Content-Type", "video/mp4")
	w.Write(data)
}
/**
 * @Author:lemon
 * @Desc:播放视频
 */
func main() {
	http.Handle("/", new(viewHandler))
	err := http.ListenAndServe(consts.VideoPlayServicePort, nil)
	if err != nil {
		fmt.Println("视频流服务启动失败")
		return
	}

}