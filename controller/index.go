package controller

import (
	"gin-demo/common"
	"gin-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage(c *gin.Context) {
	var articleList []model.ArticleCate
	var tagRow model.Tag
	var cateRow model.Category

	var cateList []model.Category

	db := common.GetDB()
	db.Find(&cateList)

	db.Table("articles").Find(&articleList)
	for i, v := range articleList {
		db.Where("id = ?", v.CateId).First(&cateRow)
		db.Where("id = ?", v.TagId).First(&tagRow)
		articleList[i].CateName = cateRow.CateName
		articleList[i].TagName = tagRow.TagName
	}

	c.HTML(http.StatusOK, "Index.html", gin.H{"articleList": articleList, "cateList": cateList})
}
