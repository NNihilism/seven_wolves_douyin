package service

import (
	"bufio"
	"context"
	"douyin/cmd/video/dao"
	"douyin/kitex_gen/video"
	"douyin/pkg/consts"
	"fmt"
	"github.com/google/uuid"
	"io/fs"
	"os"
	"path/filepath"
)

type VideoService struct {
	ctx context.Context
}
func NewVideoService(ctx context.Context)*VideoService{
	return &VideoService{ctx: ctx}
}

func (service *VideoService)GetFeed(req *video.FeedRequest)([]*dao.Video,error){
	videosDbs, err := dao.GetFeed(service.ctx, req.GetLatestTime(), req.GetToken())
	if err!=nil{
		return nil,err
	}
	return videosDbs,nil
}
func (service *VideoService)PublishVideo(req *video.PublishActionRequest)(error){
	// 从用户业务中，通过token获取userId
	var authorId int64
	authorId = 666

	// 随机生成一个文件名
	fileName := uuid.New().String()

	// 将视频流写到本地文件中，先不考虑用云服务存储 TODO:路径前缀
	file, err := os.OpenFile(filepath.Base("./"+fileName), os.O_CREATE, fs.ModeAppend)
	if err != nil {
		fmt.Println("create file error, fileName is: ", fileName)
		return  err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(req.Data)
	if err != nil {
		fmt.Println("write into server error, video's title is : ", req.Title)
		return  err
	}
	// 写成功后存入数据库
	playUrl := consts.VideoPlayUrlPrefix + fileName
	// todo 这里需要后续完善，封面的url
	coverUrl := consts.VideoCoverUrlPrefix + ""
	err = dao.CreateVideo()
	return nil
}


