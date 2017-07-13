package models

import "time"

type Mapping struct {
	ID          int64
	Name        string
	Content     string
	PublishDate time.Time
}
