package dal

import (
	"context"
	"douyin/kitex_gen/follows"
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
func UpdateFollowStatus(ctx context.Context, req *follows.RelationActionRequest) {
	actionType, _ := strconv.Atoi(req.ActionType)
	if actionType == 1 {
		follow := Follow{UserId: int(req.UserId), FollowerId: int(req.ToUserId)}
		if err := DB.Create(&follow).Error; err != nil {
			fmt.Println("关注失败", err)
		}
	} else {
		if err := DB.Where("userid = ? and followerid = ?", int(req.UserId), int(req.ToUserId)).Delete(Follow{}).Error; err != nil {
			fmt.Println("取关失败", err)
		}
	}
}
