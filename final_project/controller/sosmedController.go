package controller

import (
	"final_project/helpers"
	"final_project/models"
	"final_project/service"

	"github.com/gin-gonic/gin"
)

type SosMedController struct {
	sosMedService service.SosMedService
}

func NewSosMedController(service *service.SosMedService) *SosMedController {
	return &SosMedController{
		sosMedService: *service,
	}
}

func (s *SosMedController) CreateSosMed(c *gin.Context) {
	sosMed := models.SosMed{}

	err := c.ShouldBind(&sosMed)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	validationResult, err := helpers.ValidatePayloadSosMed(sosMed)
	if err != nil {
		helpers.InvalidPaylod(c, validationResult)
		return
	}

	id, _ := c.Get("id")
	sosMed.UserID = uint(id.(float64))

	res, err := s.sosMedService.CreateSosMed(sosMed)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	helpers.SosMedCreateSuccess(c, res)
}

func (s *SosMedController) GetSosMeds(c *gin.Context) {
	res, err := s.sosMedService.GetSosMeds()
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}

	helpers.SosMedGetSuccess(c, res)
}

func (s *SosMedController) UpdateSosMedById(c *gin.Context) {
	sosMed := models.SosMed{}

	id, isExist := c.Params.Get("socialMediaId")
	if !isExist {
		helpers.BadRequest(c, isExist)
		return
	}

	err := c.ShouldBind(&sosMed)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	cid, _ := c.Get("id")
	userId := uint(cid.(float64))

	res, err := s.sosMedService.UpdateSosMedById(userId, id, sosMed)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}

	helpers.SosMedUpdateSuccess(c, res)
}

func (s *SosMedController) DeleteSosMedById(c *gin.Context) {
	id, isExist := c.Params.Get("socialMediaId")
	if !isExist {
		helpers.BadRequest(c, isExist)
		return
	}

	cid, _ := c.Get("id")
	userId := uint(cid.(float64))

	res, err := s.sosMedService.DeleteSosMedById(userId, id)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}

	helpers.SosMedDeleteSuccess(c, res)
}
