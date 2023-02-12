package service

import (
	"context"
	"douyin/cmd/video/dao"
	"douyin/kitex_gen/video"
	"douyin/pkg/consts"
	"fmt"
	"github.com/google/uuid"
	"io/fs"
	"io/ioutil"
	"time"
)

type VideoService struct {
	ctx context.Context
}
func NewVideoService(ctx context.Context)*VideoService{
	return &VideoService{ctx: ctx}
}

func (service *VideoService)GetFeed(req *video.FeedRequest)([]*dao.Video,error){
	fmt.Println("api:GetFeed")
	videosDbs, err := dao.GetFeed(service.ctx, req.GetLatestTime(), req.GetToken())
	//TODO:拼接视频和图片URL
	if err!=nil{
		return nil,err
	}
	return videosDbs,nil
}
func (service *VideoService)PublishVideo(req *video.PublishActionRequest)(error){
	// 从用户业务中，通过token获取userId
	authorId := 666
	// 随机生成一个文件名
	fileName := uuid.New().String()
	// 将视频流写到本地文件中，先不考虑用云服务存储 TODO:路径前缀
	err := ioutil.WriteFile("./", req.GetData(), fs.ModeAppend)

	if err != nil {
		fmt.Println("create file error, fileName is: ", fileName)
		return  err
	}
	if err != nil {
		fmt.Println("write into server error, video's title is : ", req.Title)
		return  err
	}
	playUrl := consts.VideoPlayUrlPrefix + fileName
	coverUrl := consts.VideoCoverUrlPrefix + ""
	err = dao.CreateVideo(service.ctx,&dao.Video{
		Id:       int64(uuid.New().ID()),
		AuthorId: int64(authorId),
		PlayUrl: playUrl,
		CoverUrl: coverUrl,
		Title: req.GetTitle(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return err
}


