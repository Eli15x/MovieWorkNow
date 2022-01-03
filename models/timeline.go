package models

import (
)

type UserContent struct {
	UserId             string    `json:"userId,omitempty" bson:"userId,omitempty"`
	ContentId          string    `json:"contentId,omitempty" bson:"-"`
	Content            string    `json:"content,omitempty" bson:"-"`
	Data               string    `json:"data,omitempty" bson:"-"`
}
