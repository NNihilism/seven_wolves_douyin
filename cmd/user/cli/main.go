package main

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/pkg/consts"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	c, err := userservice.NewClient(
		// consts.UserServiceName,
		"kstudy",
		client.WithHostPorts(consts.UserServiceAddr),
	)
	if err != nil {
		// log.Error(err)
		fmt.Println(err)
	}

	// req := &api.Request{
	// 	Message1: "message1",
	// 	Message2: "message2",
	// }
	// resp, err := c.Concat(context.Background(), req)
	resp, err := c.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: "NNihilismmmm",
		Password: "pwd",
	}, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		// log.Error(err)
		fmt.Println(err)
	}
	fmt.Println("resp:", resp)
}
