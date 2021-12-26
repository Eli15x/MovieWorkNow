package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
  "go.mongodb.org/mongo-driver/mongo/options" // Opções para conecar com o mongo
  "github.com/Eli15x/MovieWorkNow/infrastructure"
)

func main() {
  // Echo instance
  e := echo.New()
  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Connecting to Mongo.
  	credential := options.Credential{
		Username:      config.MongodbUser,
		Password:      config.MongodbPassword,
		PasswordSet:   true,
		AuthSource:    config.MongodbDatabase,
		AuthMechanism: config.MongodbAuth,
	}

  /* if err := storage.GetInstance().Initialize(ctx, credential, "mongodb://"+config.MongodbHost+":"+config.MongodbPort,
		config.MongodbDatabase); err != nil {
		e.Logger.Fatal("[Users-Intervention] Could not resolve Data access layer. Error: ", err)
	}*/

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}