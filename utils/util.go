package utils

import (
	"github.com/satori/go.uuid"
)

func CreateCodeId() string {
	return uuid.NewV4().String()
}

func sliceToStrMap(elements []string) map[string]string {
    elementMap := make(map[string]string)
    for _, s := range elements {
        elementMap[s] = s
    }
    return elementMap
}

func sliceToIntMap(elements []string) map[string]int {
    elementMap := make(map[string]int)
    for _, s := range elements {
        elementMap[s]++
    }
    return elementMap
}

