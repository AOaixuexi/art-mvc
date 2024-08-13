package service

import (
	"article-manager/model"
	"net/http"
	_ "net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

// 创建用户
func (s *Service) CreateUser(c *gin.Context) {
	var (
		err   error
		param *model.User
	)

	if err = c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if param.Name == "" || param.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The name or password is required"})
		return
	}

	err = s.dao.CreateAUser(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Register successfully"})
}

// 验证是否有该用户
func (s *Service) FindAUser(c *gin.Context) {
	var (
		err   error
		param *model.User
	)

	if err = c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if param.Name == "" || param.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The name or password is required"})
		return
	}

	_, err = s.dao.GetAUser(param.Name, param.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The user name or password is incrrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
}
