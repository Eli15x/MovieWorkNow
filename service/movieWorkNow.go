package service

import (
	"sync"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/Eli15x/MovieWorkNow/storage"
	"github.com/Eli15x/MovieWorkNow/models"

)

var (
	instance Command
	once     sync.Once
)

type Command interface {
	CreateNewProfile(ctx echo.Context, name string,email string,password string,birthDate time.Time) error
}

type MovieWorkNowService struct{}

func GetInstance() Command {
	once.Do(func() {
		instance = &MovieWorkNowService{}
	})
	return instance
}

func (m *MovieWorkNowService)CreateNewProfile(ctx echo.Context,name string, email string, password string, birthDate time.Time) error {
	profile := models.Profile {
		UserId: "1223",
		Name : name,
		Email: email,
		PassWord: password,
		BirthDate: birthDate,
	}

	_, err := storage.GetInstance().Insert(ctx,"profile",profile)
	if err != nil {
		return ctx.String(403,"Create New Profile: problem to insert into MongoDB")
	}

	return  nil
}