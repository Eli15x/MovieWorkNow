package models

import (
	"time"
)

type Content struct {
	ContentId          string     `json:"contentId,omitempty" bson:"contentId,omitempty"`
	UserId             string     `json:"userId,omitempty" bson:"userId,omitempty"`
	Content            string     `json:"content,omitempty" bson:"-"`
	Data               time.Time  `json:"data,omitempty" bson:"-"`
}

