package handlers

import (
	"net/http"
    "strings"
	"github.com/Eli15x/MovieWorkNow/src/service"
	"github.com/Eli15x/MovieWorkNow/src/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CreateProfile(c echo.Context) error {

	name := c.Param("name")
	email := c.Param("email")
	password := c.Param("password")

	if name == "" {
		return c.String(400,"Create Profile Error: name not find")
	}

	if email == "" {
		return c.String(400,"Create Profile Error: email not find")
	}

	if password == "" {
		return c.String(400,"Create Profile Error: password not find")
	}

	err := service.GetInstanceProfile().CreateNewProfile(c,name,email,password)
	if err != nil{
		return c.String(400, err.Error())
	}
	

	return c.String(http.StatusOK, "Ok")
}

func AddInformationProfile(c echo.Context) error {

	id := c.Param("id")
	job := c.Param("job")
	message := c.Param("message")


	if id == "" {
		return c.String(400,"Create Profile Error: id not find")
	}

	if job == "" {
		return c.String(400,"Create Profile Error: job not find")
	}

	if message == "" {
		return c.String(400,"Create Profile Error: message not find")
	}

	var jobList []string

	jobList = strings.Split(job, "-")

	err := service.GetInstanceProfile().AddInformationProfile(c,id,jobList,message)
	if err != nil{
		return c.String(403,err.Error())
	}

	return c.String(http.StatusOK, "Ok")
}

func GetInformationByUserIdProfile(c echo.Context) error {

	id := c.Param("id")


	if id == "" {
		return c.String(400,"Create Profile Error: id not find")
	}

	result, err := service.GetInstanceProfile().GetInformationProfile(c,id)
	if err != nil{
		return c.String(400, err.Error())
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	return c.JSON(http.StatusOK, result)	
}

func AddRelationFollow(c echo.Context) error {

	id := c.Param("id")
	companieId := c.Param("companieId")


	if id == "" {
		return c.String(400,"Create Profile Error: id not find")
	}

	if companieId == "" {
		return c.String(400,"Create Profile Error: id not find")
	}

	result, err := service.GetInstanceProfile().GetInformationProfile(c,id)
	if err != nil{
		return c.String(400, err.Error())
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	return c.JSON(http.StatusOK, result)	
}


func AddRelationFriend(c echo.Context) error {

	UserId := c.Param("userid")
	FriendId := c.Param("friendid")


	if UserId == "" {
		return c.String(400,"Create Profile Error: UserId_user not find")
	}

	if FriendId == "" {
		return c.String(400,"Create Profile Error: UserId not find")
	}

	var friend models.Friend
	err := service.GetInstanceProfile().AddRelationFriendProfile(c,UserId,FriendId,&friend)
	if err != nil{
		return c.String(403, err.Error())
	}

	err = service.GetInstanceProfile().AddRelationFriendProfile(c,FriendId,UserId,&friend)
	if err != nil{
		return c.String(403, err.Error())
	}

	return c.JSON(http.StatusOK, "ok")	
}

func AddRequestFriend(c echo.Context) error {

	UserId := c.Param("userid")
	FriendId := c.Param("friendid")


	if UserId == "" {
		return c.String(400,"Create Profile Error: UserId_user not find")
	}

	if FriendId == "" {
		return c.String(400,"Create Profile Error: UserId not find")
	}

	var friend models.Friend
	err := service.GetInstanceProfile().AddRequestFriend(c,UserId,FriendId,&friend)
	if err != nil{
		return c.String(403, err.Error())
	}

	return c.JSON(http.StatusOK, "ok")	
}



func AddContent(c echo.Context) error {

	userId := c.Param("id")
	content := c.Param("content")

	if userId == "" {
		return c.String(400,"AddContent Error: UserId not find")
	}

	if content == "" {
		return c.String(400,"AddContent Error: UserId not find")
	}

	err := service.GetInstanceProfile().AddContent(c,userId,content)
	if err != nil{
		return c.String(400, err.Error())
	}

	return c.JSON(http.StatusOK, "ok")	
}

