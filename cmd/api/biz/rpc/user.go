package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/pkg/consts"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"gorm.io/plugin/opentelemetry/provider"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	// 链路追踪相关
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)

	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		// client.WithMiddleware(mw.CommonMiddleware),
		// client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	// resp, err := userClient.CreateUser(ctx, req)
	// if err != nil {
	// 	return nil, err
	// }
	return userClient.CreateUser(ctx, req)
}

// func CreateUser(ctx context.Context, req *user.CreateUserRequest) error {
// 	resp, err := userClient.CreateUser(ctx, req)
// 	if err != nil {
// 		return err
// 	}
// 	if resp.BaseResp.StatusCode != 0 {
// 		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
// 	}
// 	return nil
// }

// // CheckUser check user info
// func CheckUser(ctx context.Context, req *demouser.CheckUserRequest) (int64, error) {
// 	resp, err := userClient.CheckUser(ctx, req)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if resp.BaseResp.StatusCode != 0 {
// 		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
// 	}
// 	return resp.UserId, nil
// }
