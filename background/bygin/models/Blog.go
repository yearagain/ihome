package models

import (
	"fmt"
	"gorm.io/gorm"
	"ihome/background/bygin/common"
)

type Blog struct {
	gorm.Model
	Title   string `gorm:"NOT NULL"`
	Content string `gorm:"NOT NULL`
}

//AddBlogDB 添加博客
func AddBlogDB(title string, blogcontent string) error {
	blog := Blog{Title: title, Content: blogcontent}
	common.Db.Create(&blog)
	var err error = nil
	return err
}

//GetBlogListDB 查询列表
func GetBlogListDB() (blog []Blog, err error) {
	e := common.Db.Order("created_at desc").Select("id", "title", "created_at").Find(&blog)
	if e != nil {
		fmt.Println(e.Error)
		err = e.Error
	} else {
		err = nil
	}
	return
}

//GetBlogDB 查询博客
func GetBlogDB(id int) (blog Blog, err error) {
	e := common.Db.First(&blog, id)
	if e != nil {
		fmt.Println(e.Error)
		err = e.Error
	} else {
		err = nil
	}
	return
}

//DelBlog 删除博客
func DelBlog(id int) bool {
	delblog := Blog{}
	delblog.ID = uint(id)
	e := common.Db.Delete(&delblog)
	if e.Error != nil {

		return false
	}
	return true
}
