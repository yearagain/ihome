package controllers

import (
	"github.com/gin-gonic/gin"
	"ihome/background/bygin/common"
	"ihome/background/bygin/models"
	"net/http"
)

func HTTPInterCeptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		token:= c.Request.FormValue("token")
		//验证登录token是否有效
		if token!=common.Token  {
			// w.WriteHeader(http.StatusForbidden)
			// token校验失败则跳转到登录页面
			c.Abort()
			resp := models.NewRespMsg(
				int(common.StatusTokenInvalid),
				"token无效，请检查登陆状态！",
				nil,
			)
			c.JSON(http.StatusUnauthorized, resp)
			return
		}
		c.Next()
	}
}
