package models

type Friend struct {
	UserId_user         string    `json:"userId_user,omitempty" bson:"UserId_user,omitempty"`
	UserId              []UserId   `json:"userId,omitempty" bson:"UserId"`
}
