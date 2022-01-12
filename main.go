package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
  "github.com/Eli15x/MovieWorkNow/src/storage"
  "github.com/Eli15x/MovieWorkNow/src/handlers"
  "context"
  "time"
)

func main() {
  // Echo instance
  e := echo.New()

  //Context
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  //Connection to Mongo
  if err := storage.GetInstance().Initialize(ctx); err != nil {
		e.Logger.Fatal("[MONGO DB - MovieWorkNow] Could not resolve Data access layer. Error: ", err)
	}

  // handler
	profile := e.Group("/profile")
	profile.GET("/name/:name/email/:email/password/:password", handlers.CreateProfile)
  profile.GET("/id/:id/job/:job/message/:message", handlers.AddInformationProfile)
  profile.GET("/id/:id", handlers.GetInformationByUserIdProfile)
  profile.GET("/userid_user/:userId_user/userId/:userId", handlers.AddRelationFriend)
  profileCompanie := e.Group("/profileCompanie")
	profileCompanie.GET("/name/:name/email/:email/password/:password", handlers.CreateProfileCompanie)
  profileCompanie.GET("/companieId/:companieId/job/:job/message/:message", handlers.AddInformationProfileCompanie)
  profileCompanie.GET("/id/:id", handlers.GetInformationByUserIdProfileCompanie)


  e.Logger.Fatal(e.Start(":1323"))
}