package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// MGetUserByNames multiple get list of user info
func (s *GetUserService) QueryUserByName(req *user.CheckUserRequest) (*user.User, error) {
	modelUsers, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return pack.User(modelUsers[0]), nil
}

// MGetUserByNames multiple get list of user info
func (s *GetUserService) MGetUserByNames(req *user.MGetUserRequest) ([]*user.User, error) {
	modelUsers, err := db.MGetUsersByNames(s.ctx, req.Usernames)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}

// MGetUser multiple get list of user info
func (s *GetUserService) MGetUser(req *user.MGetUserRequest) ([]*user.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
