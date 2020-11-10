package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"ihome/background/bygin/common"
	"ihome/background/bygin/models"
	"ihome/background/bygin/util"
	"net/http"
	"strconv"
)

//AddBlog 添加博客
func AddBlog(c *gin.Context) {
	title := c.Request.FormValue("blogtitle")
	blogcontent := c.Request.FormValue("blogcontent")
	err := models.AddBlogDB(title, blogcontent)
	if err != nil {
		fmt.Println(err)
	}
	resp := models.NewRespMsg(int(common.StatusOK), "添加成功", nil)
	data, err := json.Marshal(resp)
	util.Errhandle(err)
	c.Data(http.StatusOK, "application/json", data)
}

//GetBlogList 获取列表
func GetBlogList(c *gin.Context) {
	list, err := models.GetBlogListDB()
	if err != nil {
		fmt.Println(err)
		resp := models.NewRespMsg(int(common.StatusServerError), "未知错误！", nil)
		data, err := json.Marshal(resp)
		util.Errhandle(err)
		c.Data(http.StatusInternalServerError, "application/json", data)
		return
	}
	resp := models.NewRespMsg(int(common.StatusOK), "查询成功", list)
	data, err := json.Marshal(resp)
	util.Errhandle(err)
	c.Data(http.StatusOK, "application/json", data)
}

//GetBlogById 获取博客
func GetBlogById(c *gin.Context) {
	tmp := c.Request.FormValue("id")
	id,err:=strconv.Atoi(tmp)
	if err!=nil{
		fmt.Println(err)
	}
	blog, err := models.GetBlogDB(id)
	if err != nil {

		fmt.Println(err)
		resp := models.NewRespMsg(int(common.StatusServerError), "未知错误！", nil)
		data, err := json.Marshal(resp)
		util.Errhandle(err)
		c.Data(http.StatusInternalServerError, "application/json", data)
		return
	}
	resp := models.NewRespMsg(int(common.StatusOK), "查询成功", blog)
	data, err := json.Marshal(resp)
	util.Errhandle(err)
	c.Data(http.StatusOK, "application/json", data)
}
//DelBlog 删除博客
func DelBlog(c *gin.Context) {
	tmp := c.Request.FormValue("delblogid")
	id,err:=strconv.Atoi(tmp)
	if err!=nil{
		fmt.Println(err)
	}
	if models.DelBlog(id){
		resp := models.NewRespMsg(int(common.StatusOK), "删除成功", nil)
		data, err := json.Marshal(resp)
		util.Errhandle(err)
		c.Data(http.StatusOK, "application/json", data)
	}else{
		resp := models.NewRespMsg(int(common.StatusServerError), "删除失败", nil)
		data, err := json.Marshal(resp)
		util.Errhandle(err)
		c.Data(http.StatusInternalServerError, "application/json", data)
	}
}
