package models

type Job struct {
	JobId            string    `json:"jobId,omitempty" bson:"jobId,omitempty"`
	Name             string    `json:"name,omitempty" bson:"-"`
}
