package helpers

import (
	"final_project/models"

	"github.com/gin-gonic/gin"
)

// ERROR
func BadRequest(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(400, gin.H{
		"status":  "failed",
		"message": msg,
	})
}

// Invalid Payload
func InvalidPaylod(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(400, gin.H{
		"status":  "failed",
		"message": data,
	})
}

// Unauthorized
func Unauthorized(c *gin.Context, info string) {
	c.AbortWithStatusJSON(401, gin.H{
		"status":  "failed",
		"message": "unauthorized :" + info,
	})
}

// USER
func UserCreateSuccess(c *gin.Context, data models.User) {
	c.JSON(201, gin.H{
		"age":      data.Age,
		"email":    data.Email,
		"id":       data.ID,
		"username": data.UserName,
	})
}

func UserLoginSuccess(c *gin.Context, token string) {
	c.JSON(200, gin.H{
		"token": token,
	})
}

func UserUpdateSuccess(c *gin.Context, data models.User) {
	c.JSON(200, gin.H{
		"id":         data.ID,
		"email":      data.Email,
		"username":   data.UserName,
		"age":        data.Age,
		"updated_at": data.UpdatedAt,
	})
}

func UserDeleteSuccess(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"message": msg,
	})
}

// PHOTO
func PhotoCreateSuccess(c *gin.Context, data models.Photo) {
	c.JSON(201, gin.H{
		"id":         data.ID,
		"title":      data.Title,
		"caption":    data.Caption,
		"photo_url":  data.PhotoUrl,
		"user_id":    data.UserID,
		"created_at": data.CreatedAt,
	})
}

func PhotoGetSuccess(c *gin.Context, photos []models.Photo) {
	c.JSON(200, photos)
}

func PhotoUpdateSuccess(c *gin.Context, data models.Photo) {
	c.JSON(200, data)
}

func PhotoDeleteSuccess(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"message": msg,
	})
}

// COMMENT
func CommentCreateSuccess(c *gin.Context, data models.Comment) {
	c.JSON(201, gin.H{
		"id":         data.ID,
		"message":    data.Message,
		"photo_id":   data.PhotoID,
		"user_id":    data.UserID,
		"created_at": data.CreatedAt,
	})
}

func CommentGetSuccess(c *gin.Context, comments []models.Comment) {
	c.JSON(200, comments)
}

func CommentUpdateSuccess(c *gin.Context, data models.Comment) {
	c.JSON(200, data)
}

func CommentDeleteSuccess(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"message": msg,
	})
}

// SosMed
func SosMedCreateSuccess(c *gin.Context, data models.SosMed) {
	c.JSON(201, gin.H{
		"id":               data.ID,
		"name":             data.Name,
		"social_media_url": data.SocialMediaUrl,
		"user_id":          data.UserID,
		"created_at":       data.CreatedAt,
	})
}

func SosMedGetSuccess(c *gin.Context, sosMeds []models.SosMed) {
	c.JSON(200, sosMeds)
}

func SosMedUpdateSuccess(c *gin.Context, data models.SosMed) {
	c.JSON(200, data)
}

func SosMedDeleteSuccess(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"message": msg,
	})
}
