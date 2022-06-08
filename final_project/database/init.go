package database

import (
	"final_project/errorhandler"
	"final_project/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB(host, user, password, dbname, port string) *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	fmt.Println(config)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	fmt.Println(db, err)
	errorhandler.Check(err)

	db.AutoMigrate(models.User{}, models.Photo{}, models.SosMed{}, models.Comment{})
	return db
}

type DB struct {
	db *gorm.DB
}

type UserQuery interface {
	CreateUser(user models.User) (models.User, error)
	FindUser(where map[string]interface{}) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id interface{}) error
}

func NewUserQuery(db *gorm.DB) UserQuery {
	return &DB{db}
}

type PhotoQuery interface {
	CreatePhoto(photo models.Photo) (models.Photo, error)
	GetPhotos() ([]models.Photo, error)
	GetPhotoById(id interface{}) (models.Photo, error)
	UpdatePhotoById(photo models.Photo) (models.Photo, error)
	DeletePhotoById(id interface{}) error
}

func NewPhotoQuery(db *gorm.DB) PhotoQuery {
	return &DB{db}
}

type CommentQuery interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetComments() ([]models.Comment, error)
	GetCommentById(id interface{}) (models.Comment, error)
	UpdateCommentById(comment models.Comment) (models.Comment, error)
	DeleteCommentById(id interface{}) error
}

func NewCommentQuery(db *gorm.DB) CommentQuery {
	return &DB{db}
}

type SosMedQuery interface {
	CreateSosMed(sosMed models.SosMed) (models.SosMed, error)
	GetSosMeds() ([]models.SosMed, error)
	GetSosMedById(id interface{}) (models.SosMed, error)
	UpdateSosMedById(sosMed models.SosMed) (models.SosMed, error)
	DeleteSosMedById(id interface{}) error
}

func NewSosMedQuery(db *gorm.DB) SosMedQuery {
	return &DB{db}
}
