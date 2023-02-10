// Code generated by hertz generator.

package main

import (
	"douyin/cmd/api/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	register(h)

	rpc.InitVideoClient()

	h.Spin()
}
