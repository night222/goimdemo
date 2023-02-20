package service

import (
	"goimdemo/models"
	"goimdemo/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// UserList
// @Tags 用户
// @Summary 用户列表
// @Success 200 {string} json{"code","massage"}
// @Router /user [get]
func UserList(ctx *gin.Context) {
	data := models.GetUserList()
	ctx.JSON(200, gin.H{
		"massage": data,
	})
	return
}

// CreatUser
// @Tags 用户
// @Summary 添加用户
// @param username formData string true "用户名"
// @param password formData string true "密码"
// @param repassword formData string true "确认密码"
// @param phone formData string true "手机号"
// @Success 200 {string} json{"code","massage"}
// @Router /user [post]
func CreateUser(ctx *gin.Context) {
	user := &models.UserBasic{
		LoginTime:     time.Now(),
		HeartbeatTime: time.Now(),
		LoginOutTime:  time.Now(),
	}
	user.Name = ctx.PostForm("username")
	user.Phone = ctx.PostForm("phone")
	user.Password = ctx.PostForm("password")
	repassword := ctx.PostForm("repassword")
	if user.Password != repassword {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"massage": "password and repassword do not agree",
		})
	}
	//进行md5加密
	salt := utils.RandString(utils.ConfigData.Max)
	user.Salt = salt
	user.Password = utils.MakePassword(user.Password, salt)
	//验证手机号格式
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"massage": "创建失败，手机号格式错误!",
		})
		return
	}
	//验证用户名或者手机号是否存在
	where := "name = ? OR phone = ?"
	vailUser := models.FirstUser(where, user.Name, user.Phone)
	if vailUser.Name != "" {
		ctx.JSON(400, gin.H{
			"massage": "创建失败，用户名或手机号已存在!",
		})
		return
	}
	models.CreateUser(user)
	ctx.JSON(http.StatusCreated, gin.H{
		"massage": "创建成功",
	})
}

// DeleteUser
// @Tags 用户
// @Summary 删除用户
// @param id query string true "id"
// @Success 200 {string} json{"code","massage"}
// @Router /user [delete]
func DeleteUser(ctx *gin.Context) {
	user := &models.UserBasic{}
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 0)
	user.ID = uint(id)
	models.DeleteUser(user)
	ctx.JSON(200, gin.H{
		"message": "删除成功",
	})
}

// UpdateUser
// @Tags 用户
// @Summary 修改用户
// @param id query string true "id"
// @param username formData string false "用户名"
// @param password formData string false "密码"
// @param phone formData string false "手机号"
// @param email formData string false "邮箱"
// @Success 200 {string} json{"code","massage"}
// @Router /user [put]
func UpdateUser(ctx *gin.Context) {
	user := &models.UserBasic{}
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 0)
	user.ID = uint(id)
	user.Name = ctx.PostForm("username")
	user.Password = ctx.PostForm("password")
	user.Phone = ctx.PostForm("phone")
	user.Email = ctx.PostForm("email")
	//验证手机号和邮箱是否符合规则
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	//验证该用户是否存在
	where := "id = ?"
	vailUser := models.FirstUser(where, user.ID)
	if vailUser.Name == "" {
		ctx.JSON(400, gin.H{
			"massage": "修改失败，用户不存在!",
		})
		return
	}
	if user.Name == "" {
		user.Name = vailUser.Name
	}
	if user.Phone == "" {
		user.Phone = vailUser.Phone
	}
	where = "(name = ? OR phone = ?) AND id != ?"
	vailUser = models.FirstUser(where, user.Name, user.Phone, user.ID)
	if vailUser.Name != "" {
		ctx.JSON(400, gin.H{
			"massage": "修改失败，该用户名或手机号以存在!",
		})
		return
	}
	//加密password
	if user.Password != "" {
		user.Password = utils.MakePassword(user.Password, vailUser.Salt)
	}
	models.UpdateUser(user)
	ctx.JSON(200, gin.H{
		"massage": "修改成功",
	})
}

// Login
// @Tags 登录
// @Summary 登录
// @param username formData string true "用户名"
// @param password formData string true "密码"
// @Success 200 {string} json{"code","massage"}
// @Router /login [post]
func Login(ctx *gin.Context) {
	user := models.UserBasic{
		LoginTime: time.Now(),
	}
	user.Name = ctx.PostForm("username")
	password := ctx.PostForm("password")
	if user.Name == "" || password == "" {
		ctx.JSON(400, gin.H{
			"message": "用户名或密码为空!",
		})
		return
	}
	where := "name = ?"
	vaildUser := models.FirstUser(where, user.Name)
	if vaildUser.Name == "" {
		ctx.JSON(400, gin.H{
			"message": "用户名或密码错误！",
		})
		return
	}
	//验证密码
	if !utils.ValidPassword(password, vaildUser.Password, vaildUser.Salt) {
		ctx.JSON(400, gin.H{
			"message": "用户名或密码错误!",
		})
		return
	}
	user.ID = vaildUser.ID
	user.IsLogout = true
	//获取ip
	user.ClientIp = ctx.ClientIP()
	//获取设备信息
	user.DeviceInfo = ctx.Request.Header.Get("User-Agent")
	if user.DeviceInfo != "" {
		deviceInfos := utils.RegexpString("^Mozilla/\\d\\.\\d\\s+\\(+.+?\\)", user.DeviceInfo)
		if deviceInfos == nil {
			deviceInfos = utils.RegexpString("\\(+.+?\\)", user.DeviceInfo)
		}
		if deviceInfos != nil {
			user.DeviceInfo = deviceInfos[0]
		}
	}
	models.UpdateUser(&user)
	ctx.JSON(200, gin.H{
		"message": "登录成功！",
	})
}
