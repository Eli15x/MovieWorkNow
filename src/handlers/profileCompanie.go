package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/Eli15x/MovieWorkNow/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func CreateProfileCompanie(c *gin.Context) {

	name := c.Param("name")
	email := c.Param("email")
	password := c.Param("password")
	//birthDate, err := time.Parse("YYYY-MM-DD",c.Param("birthDate"))

	/*if err != nil {
		return c.String(403,"Create Profile Error: data format")
	}*/

	if name == "" {
		c.String(403, "Create Profile Companie Error: name not find")
		return
	}

	if email == "" {
		c.String(403, "Create Profile Companie Error: email not find")
		return
	}

	if password == "" {
		c.String(403, "Create Profile Companie Error: password not find")
		return
	}

	/*if birthDate == "" {
		return c.String(403,"Create Profile Error: birthDate not find")
	}*/

	err := service.GetInstanceProfileCompanie().CreateNewProfileCompanie(context.Background(), name, email, password)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	c.String(http.StatusOK, "Ok")
}

func AddInformationProfileCompanie(c *gin.Context) {

	companieId := c.Param("companieId")
	job := c.Param("job")
	message := c.Param("message")

	if companieId == "" {
		c.String(403, "Add information Companie Error: CompanieId not find")
		return
	}

	if job == "" {
		c.String(403, "Add information Companie Error: job not find")
		return
	}

	if message == "" {
		c.String(403, "Add information Companie Error: message not find")
		return
	}
	var jobList []string

	jobList = strings.Split(job, "-")

	err := service.GetInstanceProfileCompanie().AddInformationProfileCompanie(context.Background(), companieId, jobList, message)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	c.String(http.StatusOK, "Ok")
}

func GetInformationByUserIdProfileCompanie(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.String(403, "Create Profile Companie Error: id not find")
		return
	}

	result, err := service.GetInstanceProfileCompanie().GetInformationProfileCompanie(context.Background(), id)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}
