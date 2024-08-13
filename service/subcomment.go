package service

import (
	"article-manager/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateSubComment handles creating a new sub-comment
func (s *Service) CreateSubComment(c *gin.Context) {
	var (
		err   error
		param *model.SubComment
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

	commentIDStr := c.Param("comment_id")
	commentID, err := primitive.ObjectIDFromHex(commentIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment_id"})
		return
	}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	param.CommentID = commentID
	param.UserID = userID
	param.UserName = user.Name
	param.CreatedAt = time.Now()

	err = s.dao.CreateASubComment(c, param)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, param)

}

// GetSubComments handles retrieving sub-comments by comment ID
func (s *Service) GetSubComments(c *gin.Context) {
	commentIDStr := c.Param("comment_id")
	commentID, err := primitive.ObjectIDFromHex(commentIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment_id"})
		return
	}

	subComments, err := s.dao.GetSubCommentsByCommentID(commentID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, subComments)
}

// DeleteSubComment handles deleting a sub-comment
func (s *Service) DeleteSubComment(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user_id"})
		return
	}

	subCommentIDStr := c.Param("subcomment_id")
	subCommentID, err := primitive.ObjectIDFromHex(subCommentIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid subcomment_id"})
		return
	}

	subComment, err := s.dao.GetSubCommentByID(subCommentID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Sub-comment not found"})
		return
	}

	if subComment.UserID != userID {
		c.JSON(400, gin.H{"error": "You can only delete your own sub-comments"})
		return
	}

	err = s.dao.DeleteSubComment(c, subCommentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Sub-comment deleted successfully"})
}
