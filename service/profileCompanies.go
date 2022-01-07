package service

import (
	"sync"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/Eli15x/MovieWorkNow/storage"
	"github.com/Eli15x/MovieWorkNow/models"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/fatih/structs"
)

var (
	instance Command
	once     sync.Once
)

type Command interface {
	CreateNewProfile(ctx echo.Context, name string,email string,password string) error
	AddInformationProfile(ctx echo.Context,id string,job string, message string) error
}

type MovieWorkNowService struct{}

func GetInstance() Command {
	once.Do(func() {
		instance = &MovieWorkNowService{}
	})
	return instance
}

func (m *MovieWorkNowService)CreateNewProfileCompanie(ctx echo.Context,name string, email string, password string) error {
	profile := &models.Profile {
		UserId: "1223",
		Name : name,
		Email: email,
		PassWord: password,
	}

	profileInsert := structs.Map(profile)
	

	_, err := storage.GetInstance().Insert(ctx,"profile",profileInsert)
	if err != nil {
		return ctx.String(403,"Create New Profile Companie: problem to insert into MongoDB")
	}

	return  nil
}

func (m *MovieWorkNowService)AddInformationProfileCompanie(ctx echo.Context,id string,job string, message string) error {
	var profile models.Profile

	CompanieId := map[string]interface{}{"CompanieId": id}

	//existe com aquele id
	mgoErr := storage.GetInstance().FindOne(ctx, "profileCompanie",CompanieId, &profile)
	if mgoErr != nil {
		return ctx.String(403,"Add Information Profile Companie: problem to Find CompanieId into MongoDB")
	}

	profileUpdate := map[string]interface{} {
		"Job": job,
		"ProfileMessage":message,
	}

	fmt.Println(profileUpdate)

	change := bson.M{"$set": profileUpdate}

	_, err := storage.GetInstance().UpdateOne(ctx,"profileCompanie",CompanieId,change)
	if err != nil {
		return ctx.String(403,"Add information Profile Companie: problem to update into MongoDB")
	}

	return  nil
}