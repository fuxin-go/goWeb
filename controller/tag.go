package controller

import (
	"gin-demo/common"
	"gin-demo/model"
	"gin-demo/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func AddTag(c *gin.Context) {
	db := common.GetDB()
	var reqTag model.Tag

	c.Bind(&reqTag)

	tagName := reqTag.TagName

	if len(tagName) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "名称不能为空", gin.H{})
		return
	}
	if isTagExist(db, tagName) {

		response.Response(c, http.StatusUnprocessableEntity, 422, "当前标签已经存在", gin.H{})
		return
	}

	var newTag = model.Tag{
		TagName: tagName,
	}
	db.Create(&newTag)

	//创建完后查询数据
	var tagList []model.Tag

	db.Find(&tagList)
	response.Success(c, gin.H{}, "添加成功")

}

func EditTag(c *gin.Context) {
	db := common.GetDB()
	var reqTag model.Tag

	c.Bind(&reqTag)

	tagName := reqTag.TagName

	if len(tagName) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "名称不能为空", gin.H{})
		return
	}
	db.Model(&reqTag).Update("tag_name", tagName)
	response.Success(c, gin.H{}, "更新成功")
}

func DeleteTag(c *gin.Context) {
	db := common.GetDB()
	var id = c.Query("id")

	var tagRow model.Tag

	db.Where("id = ?", id).Delete(&tagRow)

	c.Redirect(http.StatusMovedPermanently, "/tag/list")

}

func TagEditPage(c *gin.Context) {
	db := common.GetDB()
	var id = c.Query("id")

	var tagRow model.Tag

	db.Where("id = ?", id).First(&tagRow)

	c.HTML(http.StatusOK, "TagAdd.html", gin.H{"tagRow": tagRow})
}

func TagListPage(c *gin.Context) {
	db := common.GetDB()
	var tagList []model.Tag

	db.Find(&tagList)

	c.HTML(http.StatusOK, "TagList.html", gin.H{"tagList": tagList})
}

func TagAddPage(c *gin.Context) {

	c.HTML(http.StatusOK, "TagAdd.html", gin.H{})
}

func isTagExist(db *gorm.DB, tagName string) bool {
	var tag model.Tag
	db.Where("tag_name = ?", tagName).First(&tag)
	if tag.Id != 0 {
		return true
	}
	return false
}
