package models

import (
	"time"
)

type Profile	 struct {
	UserId             string    `json:"UserId,omitempty" bson:"UserId,omitempty"`
	Name               string    `json:"Name,omitempty" bson:"Name"`
	Email              string    `json:"Email,omitempty" bson:"Email"`
	PassWord           string    `json:"PassWord,omitempty" bson:"PassWord"`
	ProfileMessage     string    `json:"ProfileMessage,omitempty" bson:"ProfileMessage"`
	Experience         string    `json:"Experience,omitempty" bson:"Experience"`
	BirthDate          time.Time `json:"Birthdate,omitempty" bson:"Birthdate"`
	Job                string    `json:"Job,omitempty" bson:"Job"` 
}