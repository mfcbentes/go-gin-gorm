package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfcbentes/go-gin-gorm/helper"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Starting server...")
	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome home")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
