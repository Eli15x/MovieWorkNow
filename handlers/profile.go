package handlers

import (
	"net/http"
	"fmt"

	"github.com/Eli15x/MovieWorkNow/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CreateProfile(c echo.Context) error {

	name := c.Param("name")
	email := c.Param("email")
	password := c.Param("password")
	fmt.Println("entrou")

	if name == "" {
		return c.String(403,"Create Profile Error: name not find")
	}

	if email == "" {
		return c.String(403,"Create Profile Error: email not find")
	}

	if password == "" {
		return c.String(403,"Create Profile Error: password not find")
	}

	err := service.GetInstanceProfile().CreateNewProfile(c,name,email,password)
	if err != nil{
		return c.String(403,"Create Profile error: error in service")
	}

	return c.String(http.StatusOK, "Ok")
}

func AddInformation(c echo.Context) error {

	id := c.Param("id")
	job := c.Param("job")
	message := c.Param("message")


	if id == "" {
		return c.String(403,"Create Profile Error: id not find")
	}

	if job == "" {
		return c.String(403,"Create Profile Error: job not find")
	}

	if message == "" {
		return c.String(403,"Create Profile Error: message not find")
	}

	err := service.GetInstanceProfile().AddInformationProfile(c,id,job,message)
	if err != nil{
		return c.String(403,"Create Profile error: error in service")
	}

	return c.String(http.StatusOK, "Ok")
}

func GetInformationByUserId(c echo.Context) error {

	id := c.Param("id")


	if id == "" {
		return c.String(403,"Create Profile Error: id not find")
	}

	profile, err := service.GetInstanceProfile().GetInformation(c,id)
	if err != nil{
		return c.String(403,"Create Profile error: error in service")
	}

	log.Infof("[GetInformation] UserId: %s Email: %s Name: %t "+
		"test: %+v\n", profile.UserId, profile.Email, profile.Name, "")

	return c.JSON(http.StatusOK, profile)	
}







