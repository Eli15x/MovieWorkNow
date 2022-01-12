package handlers

import (
	"net/http"
	"strings"
	
	"github.com/Eli15x/MovieWorkNow/src/service"
	"github.com/labstack/gommon/log"
	"github.com/labstack/echo/v4"
)

func CreateProfileCompanie(c echo.Context) error {

	name := c.Param("name")
	email := c.Param("email")
	password := c.Param("password")
	//birthDate, err := time.Parse("YYYY-MM-DD",c.Param("birthDate"))

	/*if err != nil {
		return c.String(403,"Create Profile Error: data format")
	}*/

	if name == "" {
		return c.String(403,"Create Profile Companie Error: name not find")
	}

	if email == "" {
		return c.String(403,"Create Profile Companie Error: email not find")
	}

	if password == "" {
		return c.String(403,"Create Profile Companie Error: password not find")
	}

	/*if birthDate == "" {
		return c.String(403,"Create Profile Error: birthDate not find")
	}*/

	err := service.GetInstanceProfileCompanie().CreateNewProfileCompanie(c,name,email,password)
	if err != nil{
		return c.String(403,"Create Profile Companie error: error in service")
	}

	return c.String(http.StatusOK, "Ok")
}

func AddInformationProfileCompanie(c echo.Context) error {

	companieId := c.Param("companieId")
	job := c.Param("job")
	message := c.Param("message")


	if companieId == "" {
		return c.String(403,"Add information Companie Error: CompanieId not find")
	}

	if job == "" {
		return c.String(403,"Add information Companie Error: job not find")
	}

	if message == "" {
		return c.String(403,"Add information Companie Error: message not find")
	}
	var jobList []string

	jobList = strings.Split(job, "-")

	err := service.GetInstanceProfileCompanie().AddInformationProfileCompanie(c,companieId,jobList,message)
	if err != nil{
		return c.String(403,"Add Information Profile Companie error: error in service")
	}

	return c.String(http.StatusOK, "Ok")
}

func GetInformationByUserIdProfileCompanie(c echo.Context) error {

	id := c.Param("id")


	if id == "" {
		return c.String(403,"Create Profile Companie Error: id not find")
	}

	result, err := service.GetInstanceProfileCompanie().GetInformationProfileCompanie(c,id)
	if err != nil{
		return c.String(403,"Create Profile Companie error: error in service")
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	return c.JSON(http.StatusOK, result)	
}

