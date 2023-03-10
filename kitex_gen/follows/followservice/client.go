// Code generated by Kitex v0.4.4. DO NOT EDIT.

package followservice

import (
	"context"
	follows "douyin/kitex_gen/follows"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, req *follows.RelationActionRequest, callOptions ...callopt.Option) (r *follows.RelationActionResponse, err error)
	GetFollowList(ctx context.Context, req *follows.GetFollowListRequest, callOptions ...callopt.Option) (r *follows.GetFollowListResponse, err error)
	GetFollowerList(ctx context.Context, req *follows.GetFollowerListRequest, callOptions ...callopt.Option) (r *follows.GetFollowerListResponse, err error)
	GetFollowCount(ctx context.Context, req *follows.GetFollowCountRequest, callOptions ...callopt.Option) (r *follows.GetFollowCountResponse, err error)
	GetFollowerCount(ctx context.Context, req *follows.GetFollowerCountRequest, callOptions ...callopt.Option) (r *follows.GetFollowerCountResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFollowServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFollowServiceClient struct {
	*kClient
}

func (p *kFollowServiceClient) RelationAction(ctx context.Context, req *follows.RelationActionRequest, callOptions ...callopt.Option) (r *follows.RelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, req)
}

func (p *kFollowServiceClient) GetFollowList(ctx context.Context, req *follows.GetFollowListRequest, callOptions ...callopt.Option) (r *follows.GetFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowList(ctx, req)
}

func (p *kFollowServiceClient) GetFollowerList(ctx context.Context, req *follows.GetFollowerListRequest, callOptions ...callopt.Option) (r *follows.GetFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerList(ctx, req)
}

func (p *kFollowServiceClient) GetFollowCount(ctx context.Context, req *follows.GetFollowCountRequest, callOptions ...callopt.Option) (r *follows.GetFollowCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowCount(ctx, req)
}

func (p *kFollowServiceClient) GetFollowerCount(ctx context.Context, req *follows.GetFollowerCountRequest, callOptions ...callopt.Option) (r *follows.GetFollowerCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerCount(ctx, req)
}
