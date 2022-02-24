package controller

import "C"
import (
	"gin-demo/common"
	"gin-demo/model"
	"gin-demo/response"
	"gin-demo/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register 用户注册
func Register(c *gin.Context) {
	// 获取db
	db := common.GetDB()

	reqUser := model.User{}
	c.Bind(&reqUser)

	username := reqUser.Username
	mobile := reqUser.Mobile
	password := reqUser.Password

	//开始数据验证（简单验证）
	if len(mobile) != 11 {
		/*
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":422,
				"msg":"手机号不为11位",
			})*/
		response.Response(c, http.StatusUnprocessableEntity, 422, "手机号应为11位", gin.H{})
		return
	}
	//密码验证
	if len(password) < 6 {
		/*
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":422,
				"msg":"密码必须大于6位",
			})*/
		response.Response(c, http.StatusUnprocessableEntity, 422, "密码必须大于6位", gin.H{})
		return
	}
	//如果用户名为空,生成随机字符串
	if len(username) == 0 {
		username = util.RandomString(10)
	}

	//判断手机号是否存在
	if isMobileExist(db, mobile) {
		/*
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":422,
				"msg":"当前手机号已经注册",
			})*/
		response.Response(c, http.StatusUnprocessableEntity, 422, "当前手机号已经注册", gin.H{})
		return
	}
	//然后对密码进行加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		/*
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 500,
				"msg":  "加密错误",
			})*/
		response.Response(c, http.StatusUnprocessableEntity, 422, "加密错误", gin.H{})
		return
	}
	newUser := &model.User{
		Username: username,
		Password: string(hashPassword),
		Mobile:   mobile,
	}
	db.Create(newUser)
	/*
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"注册成功",
		})*/
	//response.Success(c,gin.H{},"注册成功")
	c.HTML(http.StatusOK, "Login.html", gin.H{"reqUser": reqUser})
}

// Login 用户的登录接口
func Login(c *gin.Context) {
	db := common.GetDB()
	var reqUser model.User
	c.Bind(&reqUser)

	mobile := reqUser.Mobile
	password := reqUser.Password

	var user model.User
	//查询数据
	db.Where("mobile = ?", mobile).First(&user)
	if user.ID == 0 {
		/*
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":422,
				"msg":"用户不存在",
			})*/
		response.Response(c, http.StatusUnprocessableEntity, 422, "用户不存在", gin.H{})
		return
	}
	//密码的对比
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		/*c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"msg":"密码错误",
		})*/
		response.Response(c, http.StatusUnprocessableEntity, 422, "用户不存在", gin.H{})
		return
	}
	/*
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"登录成功",
			"data":user,
		})*/
	//response.Success(c,gin.H{"user":user},"登录成功")
	//登录成功前先生成token
	token, err := common.GenToken(user.Username, user.ID)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, "生成token失败", gin.H{})
		return
	}
	response.Success(c, gin.H{"user": user, "token": token}, "登录成功")
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Register.html", gin.H{})
}
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Login.html", gin.H{})
}

func Logout(c *gin.Context) {
	c.HTML(http.StatusOK, "Login.html", gin.H{"msg": "logout"})
}

func GetUserInfo(c *gin.Context) {
	db := common.GetDB()
	user := model.User{}
	token := c.Query("token")
	tokenData, _ := common.ParseToken(token)
	db.Where("id = ?", tokenData.UserId).First(&user)
	response.Response(c, http.StatusOK, 200, "", gin.H{"user": user})
}

func isMobileExist(db *gorm.DB, mobile string) bool {
	var user model.User
	db.Where("mobile = ?", mobile).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
