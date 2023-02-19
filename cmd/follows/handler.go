package main

import (
	"context"
	follows "douyin/kitex_gen/follows"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// RelationAction implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) RelationAction(ctx context.Context, req *follows.RelationActionRequest) (resp *follows.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) GetFollowList(ctx context.Context, req *follows.GetFollowListRequest) (resp *follows.GetFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) GetFollowerList(ctx context.Context, req *follows.GetFollowerListRequest) (resp *follows.GetFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowCount implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) GetFollowCount(ctx context.Context, req *follows.GetFollowCountRequest) (resp *follows.GetFollowCountResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerCount implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) GetFollowerCount(ctx context.Context, req *follows.GetFollowerCountRequest) (resp *follows.GetFollowerCountResponse, err error) {
	// TODO: Your code here...
	return
}
