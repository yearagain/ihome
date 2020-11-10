package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"ihome/background/bygin/common"
	"ihome/background/bygin/models"
	"ihome/background/bygin/util"
	"math/rand"
	"net/http"
)

//Login 登陆处理
func Login(c *gin.Context) {
	//获取登陆口令
	tmp := []byte(common.LoginPassWord + common.LoginPwdSalt)
	pwd := fmt.Sprintf("%x", md5.Sum(tmp))
	password := c.Request.FormValue("password")

	if password == pwd {
		//TODO 登陆鉴权
		common.Token=fmt.Sprintf("%x", md5.Sum([]byte(string(rand.Int()))))
		//
		resp := models.NewRespMsg(int(common.StatusOK), "登陆成功", common.Token)
		data, err := json.Marshal(resp)
		util.Errhandle(err)
		c.Data(http.StatusOK, "application/json", data)

	} else {
		resp := models.NewRespMsg(int(common.StatusLoginFailed), "密码错误！", nil)
		data, err := json.Marshal(resp)
		util.Errhandle(err)
		c.Data(http.StatusUnauthorized, "application/json", data)
	}

}
