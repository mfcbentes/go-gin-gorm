package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfcbentes/go-gin-gorm/data/request"
	"github.com/mfcbentes/go-gin-gorm/data/response"
	"github.com/mfcbentes/go-gin-gorm/helper"
	"github.com/mfcbentes/go-gin-gorm/service"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{tagsService: service}
}

// Create controller
func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update controller
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.ID = id

	controller.tagsService.Update(updateTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete controller
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.tagsService.Delete(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindById controller
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagsService.FindByID(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindAll controller
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
