package utils

import (
	"github.com/satori/go.uuid"
)

func CreateCodeId() string {
	return uuid.NewV4().String()
}

