package controller

import (
	"fmt"
	"gin-demo/common"
	"gin-demo/model"
	"gin-demo/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func AddCate(c *gin.Context) {
	db := common.GetDB()
	var reqCate model.Category

	c.Bind(&reqCate)

	cateName := reqCate.CateName

	if len(cateName) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "分类名称不能为空", gin.H{})
		return
	}
	if isCateExist(db, cateName) {

		response.Response(c, http.StatusUnprocessableEntity, 422, "当前分类已经存在", gin.H{})
		return
	}

	db.Create(&reqCate)

	//创建完后查询数据
	var cateList []model.Category

	db.Find(&cateList)
	response.Success(c, gin.H{}, "添加成功")

}

func EditCate(c *gin.Context) {
	db := common.GetDB()
	var reqCate model.Category

	c.Bind(&reqCate)

	fmt.Println(reqCate.CateName)
	cateName := reqCate.CateName

	if len(cateName) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "分类名称不能为空", gin.H{})
		return
	}
	db.Model(&reqCate).Update("cate_name", cateName)
	response.Success(c, gin.H{}, "更新成功")
}

func DeleteCate(c *gin.Context) {
	db := common.GetDB()
	var id = c.Query("id")

	var cateRow model.Category

	db.Where("id = ?", id).Delete(&cateRow)

	c.Redirect(http.StatusMovedPermanently, "/cate/list")

}

func CategoryEditPage(c *gin.Context) {
	db := common.GetDB()
	var id = c.Query("id")

	var cateRow model.Category

	db.Where("id = ?", id).First(&cateRow)

	c.HTML(http.StatusOK, "CategoryAdd.html", gin.H{"cateRow": cateRow})
}

func CategoryListPage(c *gin.Context) {
	db := common.GetDB()
	var cateList []model.Category

	db.Find(&cateList)

	c.HTML(http.StatusOK, "CategoryList.html", gin.H{"cateList": cateList})
}

func CategoryAddPage(c *gin.Context) {
	c.HTML(http.StatusOK, "CategoryAdd.html", gin.H{})
}

func isCateExist(db *gorm.DB, cateName string) bool {
	var cate model.Category
	db.Where("cate_name = ?", cateName).First(&cate)
	if cate.Id != 0 {
		return true
	}
	return false
}
