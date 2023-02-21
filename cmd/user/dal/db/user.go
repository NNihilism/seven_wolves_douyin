package db

import (
	"context"
	"douyin/pkg/consts"

	"gorm.io/gorm"
)

type User struct {
	// 1: i64 id
	// 2: string username
	// 3: i64 follow_count
	// 4: i64 follower_count
	// 5: string password
	gorm.Model
	// ID            int64  `json:"id"`

	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetUsers multiple get list of user info
func MGetUsersByNames(ctx context.Context, usernames []string) ([]*User, error) {
	res := make([]*User, 0)
	if len(usernames) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("username in ?", usernames).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
