package models

import (
	"time"
)

type Profile	 struct {
	UserId             string    `json:"userId,omitempty" bson:"userId,omitempty"`
	Name               string    `json:"name,omitempty" bson:"-"`
	Email              string    `json:"email,omitempty" bson:"-"`
	PassWord           string    `json:"passWord,omitempty" bson:"-"`
	ProfileMessage     string    `json:"profileMessage,omitempty" bson:"-"`
	Experience         string    `json:"experience,omitempty" bson:"-"`
	BirthDate          time.Time `json:"birthDate,omitempty" bson:"-"`
	Job                string    `json:"job,omitempty" bson:"-"` 
}