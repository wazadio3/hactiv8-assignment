package controller

import (
	"final_project/helpers"
	"final_project/models"
	"final_project/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		userService: *service,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	user := models.User{}

	err := c.ShouldBind(&user)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	validationResult, err := helpers.ValidatePayloadUser(user)
	if err != nil {
		helpers.InvalidPaylod(c, validationResult)
		return
	}

	res, err := u.userService.CreateUser(user)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	helpers.UserCreateSuccess(c, res)
}

func (u *UserController) LoginUser(c *gin.Context) {
	user := models.User{}

	err := c.ShouldBind(&user)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	res, msg := helpers.CheckLoginPayload(user)
	if msg != "" {
		helpers.InvalidPaylod(c, res)
		return
	}

	token, err := u.userService.UserLogin(user.Email, user.Password)
	if err != nil {
		helpers.BadRequest(c, token)
		return
	}
	helpers.UserLoginSuccess(c, token)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	user := models.User{}

	err := c.ShouldBind(&user)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	res, msg := helpers.CheckUpdatePayload(user)
	if msg != "" {
		helpers.InvalidPaylod(c, res)
		return
	}

	id, _ := c.Get("id")
	fmt.Println(id)

	resp, err := u.userService.UpdateUser(user, id)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	helpers.UserUpdateSuccess(c, resp)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id, _ := c.Get("id")
	fmt.Println(id)

	resp, err := u.userService.Deleteuser(id)
	if err != nil {
		helpers.BadRequest(c, err)
		return
	}

	helpers.UserDeleteSuccess(c, resp)
}
