package service

import (
	"article-manager/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 创建文章
func (s *Service) CreatePaper(c *gin.Context) {
	var (
		err   error
		param *model.Paper
	)

	userIDStr := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user_id"})
		return
	}

	user, err := s.dao.GetUserById(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Have no this user"})
		return
	}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if param.Title == "" {
		c.JSON(400, gin.H{"error": "Require Title"})
		return
	}

	param.UserID = user.ID
	param.AuthorName = user.Name
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()

	if err := s.dao.CreateAPaper(c, param); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, param)
}

// 获取全部文章
func (s *Service) GetPaperList(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user_id"})
		return
	}

	if _, err := s.dao.GetUserById(userID); err != nil {
		c.JSON(400, gin.H{"error": "Have no this user"})
		return
	}

	// 根据 user_id 查询文章
	paperList, err := s.dao.GetPapersByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(paperList) == 0 {
		c.JSON(201, gin.H{"message": "No papers found"})
		return
	}
	c.JSON(200, paperList)
}

// 获取一篇文章
func (s *Service) GetPaper(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user_id"})
		return
	}

	param := c.Param("param")
	if param == "" {
		c.JSON(400, gin.H{"error": "Parameter is required"})
		return
	}

	// 判断参数是否是ID
	if paperID, err := primitive.ObjectIDFromHex(param); err == nil {
		paper, err := s.dao.GetAPaperByIDAndUserID(paperID, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, paper)
	} else {
		papers, err := s.dao.GetPapersByTitleAndUserID(param, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, papers)
	}
}

// 更新文章
func (s *Service) UpdateAPaper(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user_id"})
		return
	}

	idStr := c.Param("id")
	paperID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	// 获取文章并验证用户是否是文章所有者
	paper, err := s.dao.GetAPaperByIDAndUserID(paperID, userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Paper not found"})
		return
	}

	if err := c.ShouldBindJSON(&paper); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	paper.UpdatedAt = time.Now()

	// 更新文章
	err = s.dao.UpdateAPaper(c, paper)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, paper)
}

// 删除文章
func (s *Service) DeleteAPaper(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user_id"})
		return
	}

	idStr := c.Param("id")
	paperID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	// 验证用户是否是文章所有者
	_, err = s.dao.GetAPaperByIDAndUserID(paperID, userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Paper not found"})
		return
	}

	// 删除文章
	err = s.dao.DeleteAPaper(c, paperID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted"})
}

// 获取文章
func (s *Service) GetPapers(c *gin.Context) {
	userName := c.Query("user_name")
	title := c.Query("paper_name")

	if userName == "" && title == "" {
		c.JSON(400, gin.H{"error": "Require user_name or paper_name"})
		return
	}

	var papers []*model.Paper
	var err error

	if userName != "" && title != "" {
		user, err := s.dao.GetUserByName(userName)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		papers, err = s.dao.GetPapersByTitleAndUserID(title, user.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, papers)

	} else if userName != "" {
		user, err := s.dao.GetUserByName(userName)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		papers, err = s.dao.GetPapersByUserID(user.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, papers)

	} else {
		papers, err = s.dao.GetPapersByTitle(title)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, papers)
	}
}
