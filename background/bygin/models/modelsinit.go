package models

import "ihome/background/bygin/common"

func init(){

	common.Db.AutoMigrate(&Blog{})
}