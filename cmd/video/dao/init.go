package dao

import (
	"douyin/pkg/consts"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"time"
)

var DB *gorm.DB
func Init(){
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	DB,err = gorm.Open(mysql.Open(consts.VideoMySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt: true,
			Logger: gormlogrus,
		},
	)
	if err != nil{
		panic(err)
	}
	fmt.Printf("数据库连接成功%v",DB)
}

