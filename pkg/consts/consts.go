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
	VideoServerHost     = "127.0.0.1"
	VideoServerPort     = ":8083"
	VideoPlayUrlPrefix  = "http://" + BaseIP + VideoAndImageServicePort + "/douyin/video/"
	VideoCoverUrlPrefix = "http://" + BaseIP + VideoAndImageServicePort + "/douyin/image/"
	/*
		@DESC:video const
		@Author:lemon
	*/
	BaseIP                   = "127.0.0.1"
	VideoStorePathPrefix     = "./cmd/video_play/static/video/" //视频存储的位置
	ImageStorePathPrefix     = "./cmd/video_play/static/image/" //图片存储的位置
	VideoMySQLDefaultDSN     = VideoDBUserName + VideoDBPassword + "@tcp(" + VideoDBAddress + ")/" + VideoDBName + "?charset=utf8&parseTime=True&loc=Local"
	VideoAndImageServicePort = ":8089"
	VideoServerName          = "video_module"

	VideoDBUserName = "seven_wolve:"
	VideoDBPassword = "a123123"
	VideoDBAddress  = "43.138.245.151:3306"
	VideoDBName     = "douyin"
)
