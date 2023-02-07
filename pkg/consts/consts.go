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
	VideoDBUserName = "root"
	VideoDBPassword = "hjh123123"
	VideoDBAddress  = "192.168.75.100:3306"
	VideoDBName     = "douyin"

	VideoTableName  = "video"
	VideoDefaultDSN = VideoDBUserName + ":" + VideoDBPassword + "@tcp(" + VideoDBAddress + ")/" + VideoDBName + "?charset=utf8&parseTime=True&loc=Local"
	//
)
