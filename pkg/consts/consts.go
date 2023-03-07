package consts

const (
	TCP         = "tcp"
	SecretKey   = "secret key"
	IdentityKey = "id"

	RedisAddr = ":6379"
	// server name
	UserServiceName = "user"
	ApiServiceName  = "api"

	//
	ExportEndpoint = ":4317"

	// ECTD server register and find
	ETCDAddress = "127.0.0.1:2379"

	// user module
	UserDefaultDSN  = "root:abc123@tcp(172.23.112.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	UserTableName   = "user"
	UserServiceAddr = ":9000"
	// video module
	VideoServerHost = "127.0.0.1"
	VideoServerPort = ":8083"
	VideoDBUserName = "root"
	VideoDBPassword = "hjh123123"
	VideoDBAddress  = "192.168.75.100:3306"
	VideoDBName     = "douyin"
	VideoTableName  = "video"
	VideoDefaultDSN = VideoDBUserName + ":" + VideoDBPassword + "@tcp(" + VideoDBAddress + ")/" + VideoDBName + "?charset=utf8&parseTime=True&loc=Local"

	VideoPlayUrlPrefix  = "http://" + BaseIP + VideoAndImageServicePort + "/douyin/video/"
	VideoCoverUrlPrefix = "http://" + BaseIP + VideoAndImageServicePort + "/douyin/image/"
	/*
		@DESC:video const
		@Author:lemon
	*/
	BaseIP                   = "127.0.0.1"
	VideoStorePathPrefix     = "./cmd/video_play/static/video/" //视频存储的位置
	ImageStorePathPrefix     = "./cmd/video_play/static/image/" //图片存储的位置
	VideoMySQLDefaultDSN     = "root:A1548223199@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	VideoAndImageServicePort = ":8089"
	VideoServerName          = "video_module"

	// follows_module
	FollowsDBUserName  = "root"
	FollowsPassword    = "abc123"
	FollowserverDBHost = "regfsdfg1"
	FollowserverPort   = "3306"
	FollowsDBName      = "douyin"
	FollowsDNS         = FollowsDBUserName + ":" + FollowsPassword + "@tcp(" + FollowserverDBHost + ":" + FollowserverPort + ")/" + FollowsDBName + "?charset=utf8&parseTime=True&loc=Local"
	FollowTableName    = "follow"
	FollowServiceName  = "follows"
	FollowServicePort  = ":8084"
	FollowserverHost   = ""
)
