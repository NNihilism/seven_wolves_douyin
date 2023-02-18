package dal

import "douyin/cmd/user/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
