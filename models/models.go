package models

import "time"

type DataVersion struct {
	Id         int64
	Hash       string
	SchemaType string
}

type Mapping struct {
	Name        string
	Content     string
	PublishDate time.Time
}
