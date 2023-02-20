package service

import (
	"goimdemo/models"
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
// @param passworld formData string true "密码"
// @param repassworld formData string true "确认密码"
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
	user.Passworld = ctx.PostForm("passworld")
	repassworld := ctx.PostForm("repassworld")
	if user.Passworld != repassworld {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"massage": "passworld and repassworld do not agree",
		})
	}
	//验证手机号格式
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"massage": "创建失败，手机号格式错误!",
		})
		return
	}
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
// @param passworld formData string false "密码"
// @param phone formData string false "手机号"
// @param email formData string false "邮箱"
// @Success 200 {string} json{"code","massage"}
// @Router /user [put]
func UpdateUser(ctx *gin.Context) {
	user := &models.UserBasic{}
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 0)
	user.ID = uint(id)
	user.Name = ctx.PostForm("username")
	user.Passworld = ctx.PostForm("passworld")
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
	models.UpdateUser(user)
	ctx.JSON(200, gin.H{
		"massage": "修改成功",
	})
}
