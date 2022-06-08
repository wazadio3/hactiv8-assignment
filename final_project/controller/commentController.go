package controller

import (
	"final_project/helpers"
	"final_project/models"
	"final_project/service"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService service.CommentService
}

func NewCommentController(service *service.CommentService) *CommentController {
	return &CommentController{
		commentService: *service,
	}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	comment := models.Comment{}

	err := ctx.ShouldBind(&comment)
	if err != nil {
		helpers.BadRequest(ctx, err)
		return
	}

	validationResult, err := helpers.ValidatePayloadComment(comment)
	if err != nil {
		helpers.InvalidPaylod(ctx, validationResult)
		return
	}

	id, _ := ctx.Get("id")
	comment.UserID = uint(id.(float64))

	res, err := c.commentService.CreateComment(comment)
	if err != nil {
		helpers.BadRequest(ctx, err)
		return
	}

	helpers.CommentCreateSuccess(ctx, res)
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	res, err := c.commentService.GetComments()
	if err != nil {
		helpers.BadRequest(ctx, err.Error())
		return
	}

	helpers.CommentGetSuccess(ctx, res)
}

func (c *CommentController) UpdateCommentById(ctx *gin.Context) {
	comment := models.Comment{}

	id, isExist := ctx.Params.Get("commentId")
	if !isExist {
		helpers.BadRequest(ctx, isExist)
		return
	}

	err := ctx.ShouldBind(&comment)
	if err != nil {
		helpers.BadRequest(ctx, err.Error())
		return
	}

	cid, _ := ctx.Get("id")
	userId := uint(cid.(float64))

	res, err := c.commentService.UpdateCommentById(userId, id, comment)
	if err != nil {
		helpers.BadRequest(ctx, err.Error())
		return
	}

	helpers.CommentUpdateSuccess(ctx, res)
}

func (c *CommentController) DeleteCommentById(ctx *gin.Context) {
	id, isExist := ctx.Params.Get("commentId")
	if !isExist {
		helpers.BadRequest(ctx, isExist)
		return
	}

	cid, _ := ctx.Get("id")
	userId := uint(cid.(float64))

	res, err := c.commentService.DeleteCommentById(userId, id)
	if err != nil {
		helpers.BadRequest(ctx, err.Error())
		return
	}

	helpers.CommentDeleteSuccess(ctx, res)
}
