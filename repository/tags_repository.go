package repository

import "github.com/mfcbentes/go-gin-gorm/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindByID(tagsId int) (tags model.Tags, err error)
	FindAll() []model.Tags
}
