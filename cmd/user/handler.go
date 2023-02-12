package main

import (
	"context"
	user "douyin/kitex_gen/user"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	return &user.CreateUserResponse{
		BaseResp: &user.BaseResp{
			StatusCode:    0,
			StatusMessage: "Yeah, created successfully",
			ServiceTime:   time.Now().Unix(),
		},
		User: &user.User{
			Id:            8888,
			Username:      "Jack",
			Password:      "123456",
			FollowCount:   80,
			FollowerCount: 100,
		},
	}, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}
