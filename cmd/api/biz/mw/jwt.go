package mw

import (
	"context"
	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/consts"
	"douyin/pkg/errno"
	"net/http"
	"time"

	"github.com/cloudwego/biz-demo/easy_note/cmd/api/biz/model/demoapi"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(consts.SecretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   consts.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &demoapi.User{
				UserID: int64(claims[consts.IdentityKey].(float64)),
			}
		},
		// Its input parameter is the return value of 'Authenticator'
		// resolve input parameter
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					consts.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// Set user information at login
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req api.UserReq
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Name) == 0 || len(req.Pwd) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			resp, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{
				Username: req.Name,
				Password: req.Pwd,
			})
			// resp.BaseResp()
			return resp.UserId, err
			// return rpc.CheckUser(context.Background(), &user.CheckUserRequest{
			// 	Username: req.Name,
			// 	Password: req.Pwd,
			// })
		},
		// login success...
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			resp := &api.UserResp{
				StatusCode: int64(code),
				StatusMsg:  "success",
				UserID:     100086, // how to get the userid....
				Token:      token,
			}

			c.JSON(http.StatusOK, resp)

			// c.JSON(http.StatusOK, utils.H{
			// 	"code":   errno.Success.ErrCode,
			// 	"token":  token,
			// 	"expire": expire.Format(time.RFC3339),
			// })
		},
		// Jwt verification failed...
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusUnauthorized, utils.H{
				"code":    errno.AuthorizationFailedErr.ErrCode,
				"message": message,
			})
		},
		// Used to set the error information contained in the response when an error occurs in the jwt verification process
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
	})
}
