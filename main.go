package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mfcbentes/go-gin-gorm/config"
	"github.com/mfcbentes/go-gin-gorm/controller"
	"github.com/mfcbentes/go-gin-gorm/helper"
	"github.com/mfcbentes/go-gin-gorm/model"
	"github.com/mfcbentes/go-gin-gorm/repository"
	"github.com/mfcbentes/go-gin-gorm/router"
	"github.com/mfcbentes/go-gin-gorm/service"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Starting server...")

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()
	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
