package main

import (
	"fmt"
	"goimdemo/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:mysqldemo@tcp(192.168.101.8:3306)/godemo?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema 如果没有表就创建表
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := models.UserBasic{
		Name:          "first user",
		Passworld:     "123123",
		LoginTime:     time.Now(),
		HeartbeatTime: time.Now(),
		LoginOutTime:  time.Now(),
	}
	db.Create(&user)

	// Read
	fmt.Println(db.First(&user))
	fmt.Printf("first %#v\n", user) // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	fmt.Println(db.Model(&user).Update("Phone", "1111"))
	fmt.Printf("update %#v\n", user)
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
