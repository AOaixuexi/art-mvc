package dao

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TestDBConnection 测试数据库连接
func TestDBConnection(t *testing.T) {
	dsn := "paper-manager:123456@tcp(127.0.0.1:3306)/paper?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("连接数据库失败：%v", err)
	}

	fmt.Println("连接数据库成功")
}
