package service

import (
	"errors"
	"final_project/database"
	"final_project/models"
)

type CommentService struct {
	commentQuery database.CommentQuery
}

func NewCommentService(query database.CommentQuery) *CommentService {
	return &CommentService{
		commentQuery: query,
	}
}

func (c *CommentService) CreateComment(commment models.Comment) (models.Comment, error) {
	res, err := c.commentQuery.CreateComment(commment)

	return res, err
}

func (c *CommentService) GetComments() ([]models.Comment, error) {
	res, err := c.commentQuery.GetComments()

	return res, err
}

func (c *CommentService) UpdateCommentById(userId uint, id interface{}, data models.Comment) (models.Comment, error) {
	comment, err := c.commentQuery.GetCommentById(id)
	if err != nil {
		return comment, err
	}

	if comment.ID != userId {
		return comment, errors.New("Comment not found")
	}

	data.ID = comment.ID
	data.CreatedAt = comment.CreatedAt
	data.UserID = comment.UserID
	data.PhotoID = comment.PhotoID

	comment, err = c.commentQuery.UpdateCommentById(data)
	if err != nil {
		return comment, err
	}

	return comment, err
}

func (c *CommentService) DeleteCommentById(userId uint, id interface{}) (string, error) {
	comment, err := c.commentQuery.GetCommentById(id)
	if err != nil {
		return "", err
	}

	if comment.UserID != userId {
		return "", errors.New("Comment not found")
	}

	err = c.commentQuery.DeleteCommentById(id)
	if err != nil {
		return "", err
	}

	return "Your comment deleted successfully", nil
}
