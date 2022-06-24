package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Eli15x/MovieWorkNow/src/models"
	"github.com/Eli15x/MovieWorkNow/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func CreateProfile(c *gin.Context) {

	name := c.Param("name")
	email := c.Param("email")
	password := c.Param("password")

	if name == "" {
		c.String(http.StatusBadRequest, "Create Profile Error: name not find")
		return
	}

	if email == "" {
		c.String(400, "Create Profile Error: email not find")
		return
	}

	if password == "" {
		c.String(400, "Create Profile Error: password not find")
		return
	}

	err := service.GetInstanceProfile().CreateNewProfile(context.Background(), name, email, password)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "Ok")
}

func AddInformationProfile(c *gin.Context) {

	id := c.Param("id")
	job := c.Param("job")
	message := c.Param("message")

	if id == "" {
		c.String(400, "Create Profile Error: id not find")
		return
	}

	if job == "" {
		c.String(400, "Create Profile Error: job not find")
		return
	}

	if message == "" {
		c.String(400, "Create Profile Error: message not find")
		return
	}

	var jobList []string

	jobList = strings.Split(job, "-")

	err := service.GetInstanceProfile().AddInformationProfile(context.Background(), id, jobList, message)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	c.String(http.StatusOK, "Ok")
}

func GetInformationByUserIdProfile(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.String(400, "Create Profile Error: id not find")
		return
	}

	result, err := service.GetInstanceProfile().GetInformationProfile(context.Background(), id)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func AddRelationFollow(c *gin.Context) {

	id := c.Param("id")
	companieId := c.Param("companieId")

	if id == "" {
		c.String(400, "Create Profile Error: id not find")
		return
	}

	if companieId == "" {
		c.String(400, "Create Profile Error: id not find")
		return
	}

	result, err := service.GetInstanceProfile().GetInformationProfile(context.Background(), id)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func AddRelationFriend(c *gin.Context) {

	UserId := c.Param("userid")
	FriendId := c.Param("friendid")

	if UserId == "" {
		c.String(400, "Create Profile Error: UserId_user not find")
		return
	}

	if FriendId == "" {
		c.String(400, "Create Profile Error: UserId not find")
		return
	}

	var friend models.Friend
	err := service.GetInstanceProfile().AddRelationFriendProfile(context.Background(), UserId, FriendId, &friend)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	err = service.GetInstanceProfile().AddRelationFriendProfile(context.Background(), FriendId, UserId, &friend)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func AddRequestFriend(c *gin.Context) {

	UserId := c.Param("userid")
	FriendId := c.Param("friendid")

	if UserId == "" {
		c.String(400, "Create Profile Error: UserId_user not find")
		return
	}

	if FriendId == "" {
		c.String(400, "Create Profile Error: UserId not find")
		return
	}

	var friend models.Friend
	err := service.GetInstanceProfile().AddRequestFriend(context.Background(), UserId, FriendId, &friend)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func CheckInformation(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	email := json_map["email"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.
	password := json_map["password"].(string)

	var profile models.Profile
	if email == "" {
		c.String(400, "AddContent Error: email not find")
		return
	}

	if password == "" {
		c.String(400, "AddContent Error: password not find")
		return
	}

	profile_info, err := service.GetInstanceProfile().CheckInformationValid(context.Background(), email, password, &profile)
	if err != nil {
		fmt.Println(err)
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, profile_info)
}

func AddContent(c *gin.Context) {

	userId := c.Param("id")
	content := c.Param("content")

	if userId == "" {
		c.String(400, "AddContent Error: UserId not find")
		return
	}

	if content == "" {
		c.String(400, "AddContent Error: UserId not find")
		return
	}

	err := service.GetInstanceProfile().AddContent(context.Background(), userId, content)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
