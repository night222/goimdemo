package models

import (
	"goimdemo/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Passworld     string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"` //添加验证正则
	Email         string `valid:"email"`                      //添加验证规则
	Identity      string //唯一标识
	ClientIp      string
	ClinetPort    string //端口
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogout      bool   //是否下线
	DeviceInfo    string //设备形象
}

// 定义表名
func (table *UserBasic) TableName() string {
	return "user_basic"
}

// 查找
func GetUserList() []UserBasic {
	UserBasics := make([]UserBasic, 0, 200)
	utils.DB.Find(&UserBasics)
	return UserBasics
}

// 注册
func CreateUser(user *UserBasic) *gorm.DB {
	db := utils.DB.Create(user)
	return db
}

// 删除
func DeleteUser(user *UserBasic) *gorm.DB {
	return utils.DB.Delete(user)
}

// 修改
func UpdateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Model(user).Updates(user)
}

// 查询单个用户
func FirstUser(query interface{}, arg ...interface{}) UserBasic {
	user := &UserBasic{}
	utils.DB.Where(query, arg...).First(user)
	return *user
}
