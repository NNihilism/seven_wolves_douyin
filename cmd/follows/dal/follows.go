package dal

import (
	"context"
	follows "douyin/kitex_gen/follows"
	"douyin/pkg/consts"
	"fmt"
	"strconv"
)

type Follow struct {
	Id         int
	UserId     int
	FollowerId int // 关注的对象id
	Status     int // 关注状态，1为关注，0位取关
}

func (f *Follow) TableName() string {
	return consts.FollowTableName
}

// 关注/取关操作
func UpdateFollowStatus(ctx context.Context, req *follows.RelationActionRequest) (resp *follows.RelationActionResponse, err error) {
	resp = new(follows.RelationActionResponse)
	// resp.BaseResp = BaseResp.NewBaseResp()
	actionType, _ := strconv.Atoi(req.ActionType)
	if actionType == 1 {
		var count int64
		DB.Model(&Follow{}).Where("user_id = ? and follower_id = ?", req.UserId, req.ToUserId).Count(&count)
		fmt.Println(count)
		fmt.Println(req.UserId)
		fmt.Println(req.ToUserId)
		if count == 0 {
			follow := Follow{UserId: int(req.UserId), FollowerId: int(req.ToUserId)}
			if err = DB.Create(&follow).Error; err != nil {
				fmt.Println("关注失败", err)
			} else {
				// user.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
				baseResp := &follows.BaseResp{
					StatusCode:    0,
					StatusMessage: "关注成功",
				}
				resp.BaseResp = baseResp
				return
			}
		}
	} else {
		if err = DB.Where("userid = ? and followerid = ?", int(req.UserId), int(req.ToUserId)).Delete(Follow{}).Error; err != nil {
			fmt.Println("取关失败", err)
		} else {
			baseResp := &follows.BaseResp{
				StatusCode:    0,
				StatusMessage: "取关成功",
			}
			resp.BaseResp = baseResp
			return
		}
	}
	baseResp := &follows.BaseResp{
		StatusCode:    -1,
		StatusMessage: "关注/取关失败",
	}
	resp.BaseResp = baseResp
	// fmt.Println(resp)
	return
}

// 获取关注列表
func GetFollowList(ctx context.Context, req *follows.GetFollowListRequest) (resp *follows.GetFollowListResponse, err error) {
	resp = new(follows.GetFollowListResponse)
	userId := req.UserId
	users := make([]*follows.User, 0)
	followsId := make([]int, 0)
	fmt.Println("------------------------------")
	DB.Model(&Follow{}).Select("follower_id").Where("user_id = ?", userId).Scan(&followsId)
	fmt.Println(followsId)
	DB.Table("user").Where("id in ?", followsId).Find(&users)

	resp.BaseResp = &follows.BaseResp{
		StatusCode:    0,
		StatusMessage: "查询成功",
	}
	resp.Users = users
	return
}

// 获取粉丝列表
func GetFollowerList(ctx context.Context, req *follows.GetFollowerListRequest) (resp *follows.GetFollowerListResponse, err error) {
	resp = new(follows.GetFollowerListResponse)
	userId := req.UserId
	followersId := make([]int, 0)
	DB.Model(&Follow{}).Select("user_id").Where("follower_id = ?", userId).Scan(&followersId)
	users := make([]*follows.User, 0)
	fmt.Println(followersId)
	DB.Table("user").Where("id in ?", followersId).Find(&users)
	resp.BaseResp = &follows.BaseResp{
		StatusCode:    0,
		StatusMessage: "查询成功",
	}
	resp.Users = users
	return
}
