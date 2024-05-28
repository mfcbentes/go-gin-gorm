package service

import (
	"github.com/mfcbentes/go-gin-gorm/data/request"
	"github.com/mfcbentes/go-gin-gorm/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindByID(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
