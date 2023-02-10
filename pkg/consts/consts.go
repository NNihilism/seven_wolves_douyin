package consts

const (
	// server name
	UserServiceName = "user"
	ApiServiceName  = "api"

	//
	ExportEndpoint = ":4317"

	// ECTD server register and find
	ETCDAddress = "127.0.0.1:2379"

	// video module
	VideoServerHost = "127.0.0.1"
	VideoServerPort = ":8083"

	VideoDBUserName = "root"
	VideoDBPassword = "hjh123123"
	VideoDBAddress  = "192.168.75.100:3306"
	VideoDBName     = "douyin"
	VideoTableName  = "video"
	VideoDefaultDSN = VideoDBUserName + ":" + VideoDBPassword + "@tcp(" + VideoDBAddress + ")/" + VideoDBName + "?charset=utf8&parseTime=True&loc=Local"
	// todo 需要完善的常数定义
	VideoPlayUrlPrefix  = "接口前缀"
	VideoCoverUrlPrefix = "暂时为空，官方给的结构体定义没有它"
	/*
	@DESC:video
	@Author:lemon
	*/
	BaseIP="127.0.0.1"

	VideoMySQLDefaultDSN = "root:A1548223199@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	VideoPlayServicePort = ":8089"
	VideoServerName="video_module"
)
