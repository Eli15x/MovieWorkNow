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
	instanceProfileCompanie CommandProfileCompanie
	onceProfileCompanie    sync.Once
)

type CommandProfileCompanie interface {
	CreateNewProfileCompanie(ctx echo.Context, name string,email string,password string) error
	AddInformationProfileCompanie(ctx echo.Context,id string,job []string, message string) error
	GetInformationProfileCompanie(ctx echo.Context,id string) ([]bson.M, error)
	
}

type profileCompanie struct{}

func GetInstanceProfileCompanie() CommandProfileCompanie {
	onceProfileCompanie.Do(func() {
		instanceProfileCompanie = &profileCompanie{}
	})
	return instanceProfileCompanie
}

func (p *profileCompanie)CreateNewProfileCompanie(ctx echo.Context,name string, email string, password string) error {
	profile := &models.ProfileCompanie {
		CompanieId: "1223",
		Name : name,
		Email: email,
		PassWord: password,
	}

	profileInsert := structs.Map(profile)
	

	_, err := storage.GetInstance().Insert(ctx,"profileCompanie",profileInsert)
	if err != nil {
		return ctx.String(403,"Create New Profile Companie: problem to insert into MongoDB")
	}

	return  nil
}

func (p *profileCompanie)AddInformationProfileCompanie(ctx echo.Context,companieId string,job []string, message string) error {
	var profileCompanie models.ProfileCompanie

	CompanieId := map[string]interface{}{"CompanieId": companieId}

	//existe com aquele id
	mgoErr := storage.GetInstance().FindOne(ctx, "profileCompanie",CompanieId, &profileCompanie)
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

func (p *profileCompanie)GetInformationProfileCompanie(ctx echo.Context,id string) ([]bson.M, error){
	var profileCompanie models.ProfileCompanie

	CompanieId := map[string]interface{}{"CompanieId": id}

	//existe com aquele id
	result, err := repository.Find(ctx, "profileCompanie",CompanieId, &profileCompanie)
	if err != nil {
		return nil, ctx.String(403,"Add Information Profile Companie: problem to Find Id into MongoDB")
	}

	return result, nil
}