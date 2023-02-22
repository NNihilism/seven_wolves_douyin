// Code generated by hertz generator. DO NOT EDIT.

package Api

import (
	api "douyin/cmd/api/biz/handler/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/douyin", rootMw()...)
	user := root.Group("/user")
	user.GET("/login", append(_loginMw(), api.Login)...)
	user.POST("/register", append(_registerMw(), api.Register)...)
}
