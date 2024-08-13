package service

import (
	"article-manager/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 创建评论
func (s *Service) CreateComment(c *gin.Context) {
	var (
		err   error
		param *model.Comment
	)

	userIDStr := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user_id"})
		return
	}

	user, err := s.dao.GetUserById(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	paperIDStr := c.Param("paper_id")
	paperID, err := primitive.ObjectIDFromHex(paperIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid paper_id"})
		return
	}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	param.PaperID = paperID
	param.UserID = userID
	param.UserName = user.Name
	param.CreatedAt = time.Now()

	err = s.dao.CreateAComment(context.Background(), param)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, param)
}

// 获取评论列表
func (s *Service) GetComments(c *gin.Context) {
	paperIDStr := c.Param("paper_id")
	paperID, err := primitive.ObjectIDFromHex(paperIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid paper_id"})
		return
	}

	comments, err := s.dao.GetCommentsByPaperID(paperID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, comments)
}

// 删除评论
func (s *Service) DeleteComment(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user_id"})
		return
	}

	commentIDStr := c.Param("comment_id")
	commentID, err := primitive.ObjectIDFromHex(commentIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment_id"})
		return
	}

	comment, err := s.dao.GetCommentByID(commentID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	if comment.UserID != userID {
		c.JSON(400, gin.H{"error": "You can only delete your own comments"})
		return
	}

	err = s.dao.DeleteComment(context.Background(), commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Comment deleted successfully"})
}
