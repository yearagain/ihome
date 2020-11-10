package sysinit

import (
	_ "ihome/background/bygin/common"
	_ "ihome/background/bygin/models"
	"ihome/background/bygin/routers"
)

func init(){
	//accountinit()

	routers.RoutersInit()
}