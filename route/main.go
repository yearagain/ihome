package main

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		fmt.Println(err)
	}
	staticroute, err := cfg.GetValue("static", "staticroute")
	if err != nil {
		fmt.Println(err)
	}
	port, err := cfg.GetValue("static", "port")
	if err != nil {
		fmt.Println(err)
	}
	index, err := cfg.GetValue("static", "index")
	if err != nil {
		fmt.Println(err)
	}
	background, err := cfg.GetValue("static", "background")
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	// 使用gin插件支持跨域请求
	//router.Use(cors.Default())
	router.StaticFS("/static/", http.Dir(staticroute))

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, index)
	})
	router.GET("/bg", func(c *gin.Context) {
		c.Redirect(http.StatusFound, background)
	})

	router.Run(port)
}
