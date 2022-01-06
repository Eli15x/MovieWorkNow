package service

import (
	"sync"
	"github.com/labstack/echo/v4"
	"github.com/Eli15x/MovieWorkNow/storage"
	"github.com/Eli15x/MovieWorkNow/models"
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

func (m *MovieWorkNowService)CreateNewProfile(ctx echo.Context,name string, email string, password string) error {
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

func (m *MovieWorkNowService)AddInformationProfile(ctx echo.Context,id string,job string, message string) error {
	var profile models.Profile

	userID := map[string]interface{}{"UserId": id}
	
	mgoErr := storage.GetInstance().FindOne(ctx, "profile",userID, &profile)
	if mgoErr != nil {
		return ctx.String(403,"Add Information Profile: problem to Find by Id into MongoDB")
	}

	newProfile := &models.Profile {
		UserId: id,
		Name : profile.Name,
		Email: profile.Email,
		PassWord: profile.PassWord,
		Job: job,
		ProfileMessage:message,
	}

	profileInsert := structs.Map(newProfile)
	
    // mudar aqui que o updateOne tem mais um parametro que acho que é o codigo de qual é pra atualizar
	// que provavelmente será o id.

	_, err := storage.GetInstance().UpdateOne(ctx,"profile",userID,profileInsert)
	if err != nil {
		return ctx.String(403,"Create New Profile: problem to insert into MongoDB")
	}

	return  nil
}