package models

type Friend struct {
	UserId_user         string    `json:"userId_user,omitempty" bson:"userId_user,omitempty"`
	UserId              string    `json:"userId,omitempty" bson:"-"`
}