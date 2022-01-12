package models

type Follow struct {
	UserId                  string    `json:"userId,omitempty" bson:"userId,omitempty"`
	CompanieId              string    `json:"companieId,omitempty" bson:"-"`
}
