package models

import (
	"time"
)

type Chat struct {
	UserId             string     `json:"userId,omitempty" bson:"userId,omitempty"`
	MessageId          string     `json:"messageId,omitempty" bson:"-"`
	Data               time.Time  `json:"data,omitempty" bson:"-"`
}

