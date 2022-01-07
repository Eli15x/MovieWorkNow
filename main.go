package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
  "github.com/Eli15x/MovieWorkNow/storage"
  "github.com/Eli15x/MovieWorkNow/handlers"
  "context"
  "time"
)

func main() {
  // Echo instance
  e := echo.New()
  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

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
  profile.GET("/id/:id/job/:job/message/:message", handlers.AddInformation)
  profile := e.Group("/profileCompanies")
	profile.GET("/name/:name/email/:email/password/:password", handlers.CreateProfileCompanies)
  profile.GET("/companieId/:companieId/job/:job/message/:message", handlers.AddInformation)



  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}