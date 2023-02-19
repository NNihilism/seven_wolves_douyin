package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/pkg/consts"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress}) // For service discovery
	if err != nil {
		panic(err)
	}

	fmt.Println("r : ", r)
	// 链路追踪相关
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(consts.ApiServiceName),
	// 	provider.WithExportEndpoint(consts.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	// fmt.Println(r.etcdClient.Kvs)

	// c, err := userservice.NewClient(
	// 	consts.UserServiceName,
	// 	client.WithHostPorts(consts.UserServiceAddr),
	// )

	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(r), // client.WithHostPorts(consts.UserServiceAddr) We can also specify the server address
		client.WithMuxConnection(1),
		// client.WithMiddleware(mw.CommonMiddleware),
		// client.WithInstanceMW(mw.ClientMiddleware),
		// client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)

	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
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
