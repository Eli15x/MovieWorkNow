package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Eli15x/MovieWorkNow/src/handlers"
	"github.com/Eli15x/MovieWorkNow/src/storage"
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//bugsnag configure
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       "3ecac0ed23b7b1f4b863073135c602b8",
		ReleaseStage: "production",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/org/myapp"},
		// more configuration options
	})

	bugsnag.Notify(fmt.Errorf("Test error"))

	//Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	//Connection to Mongo
	if err := storage.GetInstance().Initialize(ctx); err != nil {
		bugsnag.Notify(fmt.Errorf("[MONGO DB - MovieWorkNow] Could not resolve Data access layer. Error:"))
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.GET("/profile/name/:name/email/:email/password/:password", handlers.CreateProfile)
	router.GET("/profile/id/:id/job/:job/message/:message", handlers.AddInformationProfile)
	router.GET("/profile/id/:id", handlers.GetInformationByUserIdProfile)
	router.POST("/profile/valid", handlers.CheckInformation)
	router.GET("/profile/userid/:userid/friendid/:friendid", handlers.AddRelationFriend)
	router.GET("/profile//id/:id/content/:content", handlers.AddContent)

	router.GET("/profileCompanie/name/:name/email/:email/password/:password", handlers.CreateProfileCompanie)
	router.GET("/profileCompanie/companieId/:companieId/job/:job/message/:message", handlers.AddInformationProfileCompanie)
	router.GET("/profileCompanie/id/:id", handlers.GetInformationByUserIdProfileCompanie)

	router.Run(":1323")
}
