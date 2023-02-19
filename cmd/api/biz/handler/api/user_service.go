// Code generated by hertz generator.

package api

import (
	"context"
	"fmt"

	api "douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/rpc"
	"douyin/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Login .
// @router /login [GET]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserResp)

	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /register [GET]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("req:", req)
	// userClient.CreateUser()
	rpcResp, err := rpc.CreateUser(ctx, &user.CreateUserRequest{
		Username: req.Name,
		Password: req.Pwd,
	})

	// program error, not business error
	if err != nil {
		c.JSON(consts.StatusInternalServerError, nil)
	}

	resp := &api.UserResp{
		StatusCode: rpcResp.BaseResp.StatusCode,
		StatusMsg:  rpcResp.BaseResp.StatusMessage,
		UserID:     rpcResp.User.Id,
		Token:      "", // token is ignored temporarily
	}

	c.JSON(consts.StatusOK, resp)
}
