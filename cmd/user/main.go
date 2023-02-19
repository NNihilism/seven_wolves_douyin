package main

import (
	"douyin/cmd/user/dal"
	user "douyin/kitex_gen/user/userservice"
	"log"
	"net"

	"douyin/pkg/consts"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func init() {
	dal.Init()
}
func main() {
	// register service
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.UserServiceAddr)
	if err != nil {
		panic(err)
	}

	// svr := user.NewServer(new(UserServiceImpl))
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		// server.WithMuxTransport(),
		// server.WithMiddleware(mw.CommonMiddleware),
		// server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
	)
	// fmt.Println(r.)
	// fmt.Printf("r: %v\n", r)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}

// package main

// import (
// 	"douyin/cmd/user/dal/db"
// 	user "douyin/kitex_gen/user/userservice"
// 	"log"
// )

// func init() {
// 	db.Init()
// }

// func main() {
// 	svr := user.NewServer(new(UserServiceImpl))

// 	err := svr.Run()

// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// }
