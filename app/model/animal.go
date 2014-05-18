package model

import (
	"time"
)

type Genus struct {
	Id   int    `json:id`
	Name string `json:name`
}

type Animal struct {
	Id       int       `json:id`
	Genus    Genus     `json:genus`
	Born_At  time.Time `json:created_at`
	Died_At  time.Time `json:died_at`
	Energy   Energy    `json:energy`
	Location Location  `json:location`
}
