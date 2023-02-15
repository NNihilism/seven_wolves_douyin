package service

import (
	"bytes"
	"context"
	"douyin/cmd/video/dao"
	"douyin/kitex_gen/video"
	"douyin/pkg/consts"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type VideoService struct {
	ctx context.Context
}

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func (service *VideoService) GetFeed(req *video.FeedRequest) ([]*dao.Video, error) {
	fmt.Println("api:GetFeed")
	videosDbs, err := dao.GetFeed(service.ctx, req.GetLatestTime())
	//TODO:拼接视频和图片URL
	// http://视频服务地址/douyin/video/日期/视频名字
	// http://视频服务地址/douyin/image/日期/视频名字
	if err != nil {
		return nil, err
	}
	return videosDbs, nil
}
func (service *VideoService) PublishVideo(req *video.PublishActionRequest) error {
	var authorId = getAuthorIdByToken(req.GetToken())
	// 随机生成一个文件名
	// 将视频流写到本地文件中，先不考虑用云服务存储
	date := time.Now()
	dateStr := date.Format("2006-01-02")
	randomStr := uuid.New().String()
	randomStr = strings.ReplaceAll(randomStr, "-", "")
	videoDirPath := consts.VideoStorePathPrefix + dateStr
	videoData := req.GetData()
	videoFilename := randomStr + ".mp4"

	err := storeVideoFile(videoDirPath, videoFilename, &videoData)
	if err != nil {
		return err
	}
	imagePathDir := consts.ImageStorePathPrefix + dateStr
	imageFilename := randomStr + ".png"
	err = storeImageFile(videoDirPath+"/"+videoFilename, imagePathDir, imageFilename, 1)
	if err != nil {
		return err
	}
	/*
		数据库url存 日期/名称
	*/
	playUrl := dateStr + "/" + videoFilename
	coverUrl := dateStr + "/" + imageFilename
	err = dao.CreateVideo(service.ctx, &dao.Video{
		Id:        int64(uuid.New().ID()),
		AuthorId:  authorId,
		PlayUrl:   playUrl,
		CoverUrl:  coverUrl,
		Title:     req.GetTitle(),
		CreatedAt: date,
		UpdatedAt: date,
	})
	return err
}
func (service *VideoService) GetPublishVideoList(req *video.PublishListRequest) ([]*dao.Video, error) {
	videoList, err := dao.GetVideoListByUserId(service.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}

// todo 等待用户服务的API替代此方法
func getAuthorIdByToken(token string) int64 {
	//TODO:
	return 0
}
func storeVideoFile(dirPath, filename string, data *[]byte) error {
	_, err := os.Stat(dirPath)
	if err != nil { //目录不存在
		err = os.MkdirAll(dirPath, os.ModeDir)
		if err != nil {
			fmt.Println("创建视频父目录失败:", err)
			return err
		}
	}
	err = os.WriteFile(dirPath+"/"+filename, *data, fs.ModeAppend)
	if err != nil {
		fmt.Println("create file error, filename is: ", filename)
		return err
	}
	return nil
}

/*
*

		@params : inputVideoPath  - 输入的视频路径
	              outputImagePath - 输出的图片路径及名称
				  frameNum        - 需要截取的视频帧
*/
func storeImageFile(inputVideoPath, outputImagePath, imageFilename string, frameNum int) (err error) {
	_, err = os.Stat(outputImagePath)
	if err != nil { //目录不存在
		err = os.MkdirAll(outputImagePath, os.ModeDir)
		if err != nil {
			log.Fatal("创建图片父目录失败:", err)
			return err
		}
	}
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(inputVideoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		fmt.Println("图片获取失败：", err)
		return err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println("图片获取失败：", err)
		return err
	}
	err = imaging.Save(img, outputImagePath+"/"+imageFilename)
	if err != nil {
		fmt.Println("图片获取失败：", err)
		return err
	}
	return nil
}
