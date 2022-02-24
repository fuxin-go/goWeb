package controller

import (
	"gin-demo/common"
	"gin-demo/model"
	"gin-demo/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArticle(c *gin.Context) {
	db := common.GetDB()

	var reqArticle model.Article

	c.Bind(&reqArticle)

	title := reqArticle.Title
	userId := reqArticle.UserId
	cateId := reqArticle.CateId

	if title == "" {
		response.Response(c, http.StatusUnprocessableEntity, 422, "标题不能为空", gin.H{})
		return
	}
	if userId <= 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "请登录", gin.H{})
		return
	}
	if cateId <= 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "分类不能为空", gin.H{})
		return
	}
	db.Create(&reqArticle)

	response.Success(c, gin.H{}, "添加成功")
}

func EditArticle(c *gin.Context) {
	db := common.GetDB()

	var reqArticle model.Article

	c.Bind(&reqArticle)

	title := reqArticle.Title
	userId := reqArticle.UserId
	cateId := reqArticle.CateId

	if title == "" {
		response.Response(c, http.StatusUnprocessableEntity, 422, "标题不能为空", gin.H{})
		return
	}
	if userId <= 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "请登录", gin.H{})
		return
	}
	if cateId <= 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "分类不能为空", gin.H{})
		return
	}
	db.Save(&reqArticle)

	response.Success(c, gin.H{}, "编辑成功")
}

func ArticleListPage(c *gin.Context) {
	var articleList []model.ArticleCate
	var tagRow model.Tag
	var cateRow model.Category

	db := common.GetDB()

	db.Table("articles").Find(&articleList)
	for i, v := range articleList {
		db.Where("id = ?", v.CateId).First(&cateRow)
		db.Where("id = ?", v.TagId).First(&tagRow)
		articleList[i].CateName = cateRow.CateName
		articleList[i].TagName = tagRow.TagName
	}
	c.HTML(http.StatusOK, "ArticleList.html", gin.H{"articleList": articleList})
}

func ArticleAddPage(c *gin.Context) {
	db := common.GetDB()

	var cateList []model.Category
	var tagList []model.Tag

	db.Find(&cateList)
	db.Find(&tagList)

	c.HTML(http.StatusOK, "ArticleAdd.html", gin.H{"cateList": cateList, "tagList": tagList})
}

func ArticleEditPage(c *gin.Context) {
	db := common.GetDB()
	var id = c.Query("id")

	var ArticleRow model.Article
	var cateList []model.Category
	var tagList []model.Tag

	db.Find(&cateList)
	db.Find(&tagList)

	db.Where("id = ?", id).First(&ArticleRow)

	c.HTML(http.StatusOK, "ArticleAdd.html", gin.H{"cateList": cateList, "tagList": tagList, "articleRow": ArticleRow})
}

func ArticleContent(c *gin.Context) {
	db := common.GetDB()
	var id = c.Query("id")

	var ArticleRow model.ArticleCate
	var cateRow model.Category

	db.Table("articles").Where("id = ?", id).First(&ArticleRow)
	db.Where("id = ?", ArticleRow.CateId).First(&cateRow)
	ArticleRow.CateName = cateRow.CateName

	c.HTML(http.StatusOK, "article.html", gin.H{"articleRow": ArticleRow})
}
