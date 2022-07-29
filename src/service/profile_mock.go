package service

import (
	"context"
	
	testifyMock "github.com/stretchr/testify/mock"
	"github.com/Eli15x/MovieWorkNow/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ProfileServiceMock struct {
	testifyMock.Mock
}

func (p *ProfileServiceMock) CreateNewProfile(ctx context.Context, name string, email string, password string) error {
	args := p.Called(ctx,name,email,password)
	return args.Error(0)
}

func (p *ProfileServiceMock) AddInformationProfile(ctx context.Context, id string, job []string, message string) error  {
	args := p.Called(ctx,id,job,message)
	return args.Error(0)
}

func (p *ProfileServiceMock) GetInformationProfile(ctx context.Context, id string) ([]bson.M, error) {
	args := p.Called(ctx, id)
	return args.Get(0).([]bson.M), args.Error(1)
}

func (p *ProfileServiceMock) CheckInformationValid(ctx context.Context, email string, password string, profile *models.Profile) (*models.Profile, error) {
	args := p.Called(ctx,email,password,profile)
	return args.Get(0).(*models.Profile), args.Error(1)
}

func (p *ProfileServiceMock) AddRelationFriendProfile(ctx context.Context, UserId string, FriendId string, friendUser *models.Friend) error{
	args := p.Called(ctx,UserId,FriendId,friendUser)
	return args.Error(0)
}

func (p *ProfileServiceMock) DeleteFriendRequest(ctx context.Context, UserId string, FriendId string, friendUser *models.Friend) error {
	args := p.Called(ctx,UserId,FriendId,friendUser)
	return args.Error(0)
}
func (p *ProfileServiceMock) CreateFriendTable(ctx context.Context, userId string) error{
	args := p.Called(ctx,userId)
	return args.Error(0)
}
func (p *ProfileServiceMock) AddContent(ctx context.Context, id string, content string) error {
	args := p.Called(ctx,id,content)
	return args.Error(0)
}
