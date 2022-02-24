package router

import (
	"gin-demo/controller"
	"gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	r.LoadHTMLGlob("template/**/*")
	//加载静态资源
	r.Static("/static", "./static")
	r.Static("/md", "./static/md")
	//首页页面
	r.POST("/getUserInfo", controller.GetUserInfo)
	r.GET("/index", controller.IndexPage)
	//注册
	r.POST("/user/register", controller.Register)
	r.GET("/user/register", controller.RegisterPage)
	//登录
	r.POST("/user/login", controller.Login)
	r.GET("/user/login", controller.LoginPage)
	r.GET("/user/logout", controller.Logout)

	//分类页面
	r.GET("/cate/list", controller.CategoryListPage)
	r.GET("/cate/add", controller.CategoryAddPage)
	r.GET("/cate/edit", controller.CategoryEditPage)
	r.GET("/cate/delete", controller.DeleteCate)

	r.POST("/cate/add", middleware.JwtAuthMiddleware(), controller.AddCate)
	r.POST("/cate/edit", controller.EditCate)

	//标签页面
	r.GET("/tag/list", controller.TagListPage)
	r.GET("/tag/add", controller.TagAddPage)
	r.GET("/tag/edit", controller.TagEditPage)
	r.GET("/tag/delete", controller.DeleteTag)

	r.POST("/tag/add", middleware.JwtAuthMiddleware(), controller.AddTag)
	r.POST("/tag/edit", controller.EditTag)

	r.GET("/article/list", controller.ArticleListPage)
	r.GET("/article/add", controller.ArticleAddPage)
	r.GET("/article/edit", controller.ArticleEditPage)

	r.POST("/article/add", middleware.JwtAuthMiddleware(), controller.AddArticle)
	r.POST("/article/edit", middleware.JwtAuthMiddleware(), controller.EditArticle)

	r.GET("/article", controller.ArticleContent)

	return r
}
