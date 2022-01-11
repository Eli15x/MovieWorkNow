package service

import (
	"sync"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/Eli15x/MovieWorkNow/storage"
	"github.com/Eli15x/MovieWorkNow/models"
	"github.com/Eli15x/MovieWorkNow/repository"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/fatih/structs"
)

var (
	instanceProfile CommandProfile
	onceProfile     sync.Once
)

type CommandProfile interface {
	CreateNewProfile(ctx echo.Context, name string,email string,password string) error
	AddInformationProfile(ctx echo.Context,id string,job []string, message string) error
	GetInformationProfile(ctx echo.Context,id string) ([]bson.M, error)
}

type profile struct{}

func GetInstanceProfile() CommandProfile {
	onceProfile.Do(func() {
		instanceProfile = &profile{}
	})
	return instanceProfile
}

func (p *profile)CreateNewProfile(ctx echo.Context,name string, email string, password string) error {
	profile := &models.Profile {
		UserId: "1223",
		Name : name,
		Email: email,
		PassWord: password,
	}

	profileInsert := structs.Map(profile)
	

	_, err := storage.GetInstance().Insert(ctx,"profile",profileInsert)
	if err != nil {
		return ctx.String(403,"Create New Profile: problem to insert into MongoDB")
	}

	return  nil
}

func (p *profile)AddInformationProfile(ctx echo.Context,id string,job []string, message string) error {
	var profile models.Profile

	userId := map[string]interface{}{"UserId": id}

	//existe com aquele id
	mgoErr := storage.GetInstance().FindOne(ctx, "profile",userId, &profile)
	if mgoErr != nil {
		return ctx.String(403,"Add Information Profile: problem to Find Id into MongoDB")
	}

	profileUpdate := map[string]interface{} {
		"Job": job,
		"ProfileMessage":message,
	}

	fmt.Println(profileUpdate)

	change := bson.M{"$set": profileUpdate}

	_, err := storage.GetInstance().UpdateOne(ctx,"profile",userId,change)
	if err != nil {
		return ctx.String(403,"Create New Profile: problem to update into MongoDB")
	}

	return  nil
}

func (p *profile)GetInformationProfile(ctx echo.Context,id string) ([]bson.M, error){
	var profile models.Profile

	userId := map[string]interface{}{"UserId": id}

	//existe com aquele id
	result, err := repository.Find(ctx, "profile",userId, &profile)
	if err != nil {
		return nil, ctx.String(403,"Add Information Profile: problem to Find Id into MongoDB")
	}

	fmt.Println(result)

	/*	profileResult := models.Profile{
		UserId: "1223",
		Name : cursor.Name,
		Email: cursor.Email,
		PassWord: cursor.PassWord,
		BirthDate: cursor.BirthDate,
		Job: cursor.Job,
		ProfileMessage:cursor.ProfileMessage,
		Experience: cursor.Experience,
	} */

	return result, nil
}


