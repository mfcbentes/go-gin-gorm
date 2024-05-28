package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/mfcbentes/go-gin-gorm/data/request"
	"github.com/mfcbentes/go-gin-gorm/data/response"
	"github.com/mfcbentes/go-gin-gorm/helper"
	"github.com/mfcbentes/go-gin-gorm/model"
	"github.com/mfcbentes/go-gin-gorm/repository"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	validate       *validator.Validate
}

func NewTagsServiceImpl(tagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{TagsRepository: tagsRepository, validate: validate}
}

// Create implements TagsService.
func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService.
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService.
func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()
	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			ID:   value.ID,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindByID implements TagsService.
func (t *TagsServiceImpl) FindByID(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindByID(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		ID:   tagData.ID,
		Name: tagData.Name,
	}

	return tagResponse
}

// Update implements TagsService.
func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindByID(tags.ID)
	helper.ErrorPanic(err)

	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}
