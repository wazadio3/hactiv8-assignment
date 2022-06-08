package controller

import (
	"final_project/helpers"
	"final_project/models"
	"final_project/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService service.PhotoService
}

func NewPhotoController(service *service.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: *service,
	}
}

func (p *PhotoController) CreatePhoto(c *gin.Context) {
	photo := models.Photo{}

	err := c.ShouldBind(&photo)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	validationResult, err := helpers.ValidatePayloadPhoto(photo)
	if err != nil {
		helpers.InvalidPaylod(c, validationResult)
		return
	}

	id, _ := c.Get("id")
	fmt.Println(id)
	photo.UserID = uint(id.(float64))

	res, err := p.photoService.CreatePhoto(photo)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	helpers.PhotoCreateSuccess(c, res)
}

func (p *PhotoController) GetPhotos(c *gin.Context) {
	id, _ := c.Get("id")
	email, _ := c.Get("email")
	username, _ := c.Get("username")
	fmt.Println(id)
	fmt.Println(email)
	fmt.Println(username)

	res, err := p.photoService.GetPhotos()
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}

	helpers.PhotoGetSuccess(c, res)
}

func (p *PhotoController) UpdatePhotoById(c *gin.Context) {
	photo := models.Photo{}

	id, isExist := c.Params.Get("photoId")
	if !isExist {
		helpers.BadRequest(c, isExist)
		return
	}

	err := c.ShouldBind(&photo)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	cid, _ := c.Get("id")
	userId := uint(cid.(float64))

	res, err := p.photoService.UpdatePhotoById(userId, id, photo)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}

	helpers.PhotoUpdateSuccess(c, res)
}

func (p *PhotoController) DeletePhotoById(c *gin.Context) {
	id, isExist := c.Params.Get("photoId")
	if !isExist {
		helpers.BadRequest(c, isExist)
		return
	}

	cid, _ := c.Get("id")
	userId := uint(cid.(float64))

	res, err := p.photoService.DeletePhotoById(userId, id)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}

	helpers.PhotoDeleteSuccess(c, res)
}
