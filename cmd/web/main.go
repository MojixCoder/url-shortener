package main

import (
	"github.com/MojixCoder/awesomeProject/pkg/config"
	"github.com/MojixCoder/awesomeProject/pkg/db"
	"github.com/MojixCoder/awesomeProject/pkg/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	appConfig := config.GetConfig()

	client, err := db.GetDb()
	if err != nil {
		log.Fatal(err)
	}

	// Set Repo for handlers
	handlers.SetRepo(handlers.NewRepo(&appConfig, client))

	// Routers
	r.POST("/link-shortener", handlers.Repo.ShortenedLinkCreate)
	r.GET("/:slug", handlers.Repo.RedirectToLink)

	// Running server
	r.Run(appConfig.Port)
}
