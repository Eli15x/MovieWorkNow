package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/bugsnag/bugsnag-go/v2"
  "github.com/Eli15x/MovieWorkNow/src/storage"
  "github.com/Eli15x/MovieWorkNow/src/handlers"
  "context"
  "time"
  "fmt"
)

func main() {
  //bugsnag configure
  bugsnag.Configure(bugsnag.Configuration{
      APIKey:          "3ecac0ed23b7b1f4b863073135c602b8",
      ReleaseStage:    "production",
      // The import paths for the Go packages containing your source files
      ProjectPackages: []string{"main", "github.com/org/myapp"},
      // more configuration options
  })

  bugsnag.Notify(fmt.Errorf("Test error"))

  // Echo instance
  e := echo.New()
  e.Use(middleware.CORS())
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  //Context
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  //Connection to Mongo
  if err := storage.GetInstance().Initialize(ctx); err != nil {
		e.Logger.Fatal("[MONGO DB - MovieWorkNow] Could not resolve Data access layer. Error: ", err)
    bugsnag.Notify(fmt.Errorf("[MONGO DB - MovieWorkNow] Could not resolve Data access layer. Error:"))
	}

  // Handler
	profile := e.Group("/profile")
	profile.GET("/name/:name/email/:email/password/:password", handlers.CreateProfile)
  profile.GET("/id/:id/job/:job/message/:message", handlers.AddInformationProfile)
  profile.GET("/id/:id", handlers.GetInformationByUserIdProfile)
  profile.POST("/valid", handlers.CheckInformation)
  profile.GET("/userid/:userid/friendid/:friendid", handlers.AddRelationFriend)
  profile.GET("/id/:id/content/:content", handlers.AddContent)
  profileCompanie := e.Group("/profileCompanie")
	profileCompanie.GET("/name/:name/email/:email/password/:password", handlers.CreateProfileCompanie)
  profileCompanie.GET("/companieId/:companieId/job/:job/message/:message", handlers.AddInformationProfileCompanie)
  profileCompanie.GET("/id/:id", handlers.GetInformationByUserIdProfileCompanie)


  e.Logger.Fatal(e.Start(":1323"))
}