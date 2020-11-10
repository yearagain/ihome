package routers

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ihome/background/bygin/controllers"
)

func RoutersInit() {
	//打开配置文件
	cfg, err := goconfig.LoadConfigFile("./conf/conf.ini")
	if err != nil {
		fmt.Println(err)
	}
	//创建gin对象
	router := gin.Default()
	//HTTP中间件
	router.Use(cors.Default()) //允许跨域请求
	//登录
	router.POST("/bglogin", controllers.Login)
	//查询
	router.POST("/getbloglist",controllers.GetBlogList)
	router.POST("/getblogbyid",controllers.GetBlogById)
	//HTTP拦截器
	router.Use(controllers.HTTPInterCeptor())
	//业务
	router.POST("/addblog", controllers.AddBlog)
	router.POST("/delblog", controllers.DelBlog)
	router.POST("/test", func(c *gin.Context) {
		//登录状态测试-无需代码
	})

	port, err := cfg.GetValue("router", "port")
	if err != nil {
		fmt.Println(err)
	}

	router.Run(port)

}
