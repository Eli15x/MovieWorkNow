package models

import (
	"time"
)

type Chat struct {
	UserId             string    `json:"userId,omitempty" bson:"userId,omitempty"`
	Name               string    `json:"name,omitempty" bson:"-"`
	Email              string    `json:"email,omitempty" bson:"-"`
	PassWord           string    `json:"passWord,omitempty" bson:"-"`
	ProfileMessage     string    `json:"profileMessage,omitempty" bson:"-"`
	Experience         string    `json:"experience,omitempty" bson:"-"`
	BirthDate          time.Time `json:"birthDate,omitempty" bson:"-"`
	Job                string    `json:"birthDate,omitempty" bson:"-"` 
	//ver depois como criar um array para job, se é por aqui mesmo ou não etc.
}