package model

type Phylum struct {
	Id   int    `json:id`
	Name string `json:name`
}

type Plant struct {
	Id       int       `json:id`
	Phylum   Phylum    `json:phylum`
	Location Location  `json:location`
	Born_At  time.Time `json:born_at`
	Died_At  time.Time `json:died_at`
	Energy   Energy    `json:energy`
}
