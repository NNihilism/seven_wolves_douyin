package db

import (
	"douyin/pkg/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB -- global connection of video module
var DB *gorm.DB

func Init() {
	var err error

	DB, err = gorm.Open(mysql.Open(consts.VideoDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		panic(err)
	}
}
