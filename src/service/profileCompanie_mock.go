package service

import (
	"context"
	
	testifyMock "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type ProfileCompanieServiceMock struct {
	testifyMock.Mock
}

func (p *ProfileCompanieServiceMock) CreateNewProfileCompanie(ctx context.Context, name string, email string, password string) error {
	args := p.Called(ctx,name,email,password)
	return args.Error(0)
}

func (p *ProfileCompanieServiceMock) AddInformationProfileCompanie(ctx context.Context, id string, job []string, message string) error  {
	args := p.Called(ctx,id,job,message)
	return args.Error(0)
}

func (p *ProfileCompanieServiceMock) GetInformationProfileCompanie(ctx context.Context, id string) ([]bson.M, error) {
	args := p.Called(ctx, id)
	return args.Get(0).([]bson.M), args.Error(1)
}
