// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	api "douyin/cmd/api/biz/router/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	api.Register(r)
}
