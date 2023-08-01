package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"oceanlearn.teach/ginessential/common"
	"oceanlearn.teach/ginessential/dto"
	"oceanlearn.teach/ginessential/model"
	"oceanlearn.teach/ginessential/response"
	"oceanlearn.teach/ginessential/util"
)

// 注册
func Register(c *gin.Context) {
	DB := common.DB
	//获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, password)

	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	//返回结果
	response.Success(c, nil, "注册成功")
}

// 登录
func Login(c *gin.Context) {
	DB := common.DB
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		//令牌生成
		log.Printf("token generate event : %v", err)
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
}

// 删除用户信息
func Delete(c *gin.Context) {
	DB := common.DB
	var user model.User
	ID := c.Param("id")
	DB.Find(&user, ID)
	if user.ID == 0 {
		//c.JSON(http.StatusNotFound, gin.H{"status":  http.StatusNotFound, "msg": "没有找到!"})
		response.Response(c, http.StatusNotFound, 500, nil, "没有找到")
		return
	}
	DB.Delete(&user)
	//c.JSON(http.StatusOK, gin.H{"status":  http.StatusOK, "msg": "删除成功"})
	response.Response(c, http.StatusOK, 200, nil, "删除成功")
}

// 查询
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func Update(c *gin.Context) {
	user, _ := c.Get("user")
	ID := c.Param("id")
	common.GetDB().Table("users").Where("id = ?", ID).Updates(gin.H{"user": dto.ToUserDto(user.(model.User))})
	c.JSON(200, gin.H{
		"msg": "修改成功",
	})
	return
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func UpdateTest(db *gorm.DB, id int64) (int64, error) {
	var user model.User
	row := db.First(&user, id)
	if row.Error == nil {
		db.Model(&user).Updates(&user)
	}
	return 0, row.Error
}
