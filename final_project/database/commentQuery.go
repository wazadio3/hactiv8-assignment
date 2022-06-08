package database

import "final_project/models"

func (c *DB) CreateComment(comment models.Comment) (models.Comment, error) {
	err := c.db.Create(&comment).Error

	return comment, err
}

func (c *DB) GetComments() ([]models.Comment, error) {
	comments := []models.Comment{}
	err := c.db.Find(&comments).Error

	return comments, err
}

func (c *DB) GetCommentById(id interface{}) (models.Comment, error) {
	comment := models.Comment{}
	err := c.db.Where("id = ?", id).First(&comment).Error

	return comment, err
}

func (c *DB) UpdateCommentById(comment models.Comment) (models.Comment, error) {
	err := c.db.Save(&comment).Error

	return comment, err
}

func (c *DB) DeleteCommentById(id interface{}) error {
	err := c.db.Delete(&models.Comment{}, id).Error

	return err
}
