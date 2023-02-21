package main

import (
	"context"
	like "douyin/kitex_gen/like"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// LikeAction implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeAction(ctx context.Context, req *like.LikeActionRequest) (resp *like.LikeActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetLikeList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) GetLikeList(ctx context.Context, req *like.GetLikeListRequest) (resp *like.GetLikeListResponse, err error) {
	// TODO: Your code here...
	return
}
