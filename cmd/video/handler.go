package main

import (
	"bufio"
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/pack"
	"douyin/cmd/video/service"
	video "douyin/kitex_gen/video"
	"douyin/pkg/consts"
	"fmt"
	"github.com/google/uuid"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	resp = new(video.FeedResponse)
	videoList, err := service.NewVideoService(ctx).GetFeed(req)
	resp.SetVideoList(pack.CovertList(videoList))
	resp.SetStatusCode(0)
	nextTime := videoList[len(videoList)-1].CreatedAt.Unix()
	resp.SetNextTime(&nextTime)
	return resp, nil
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...

	// todo: 1、鉴权

	// 从用户业务中，通过token获取userId
	var anchorId int64

	anchorId = 666

	// 随机生成一个文件名
	fileName := uuid.New().String()

	// 将视频流写到本地文件中，先不考虑用云服务存储
	file, err := os.OpenFile(filepath.Base("./"+fileName), os.O_CREATE, fs.ModeAppend)
	if err != nil {
		fmt.Println("create file error, fileName is: ", fileName)
		return &video.PublishActionResponse{StatusCode: -1}, err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(req.Data)
	if err != nil {
		fmt.Println("write into server error, video_module's title is : ", req.Title)
		return &video.PublishActionResponse{StatusCode: -1}, err
	}

	// 写成功后存入数据库
	playUrl := consts.VideoPlayUrlPrefix + fileName
	// todo 这里需要后续完善，封面的url
	coverUrl := consts.VideoCoverUrlPrefix + ""

	err = db.CreateVideo(&db.Video{
		AnchorId:    anchorId,
		PlayUrl:     playUrl,
		CoverUrl:    coverUrl,
		Title:       req.Title,
		PublishTime: time.Now(),
	})
	if err != nil {
		// 存入数据库失败，需要将文件删掉
		os.Remove(filepath.Base("./" + fileName))

		fmt.Println("write into server error, video_module's title is : ", req.Title)
		return &video.PublishActionResponse{StatusCode: -1}, err
	}

	return &video.PublishActionResponse{StatusCode: 0}, nil
}

// GetPublishVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishVideoList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
