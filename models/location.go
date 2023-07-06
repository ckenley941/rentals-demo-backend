package models

type Location struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
	Lat     int    `json:"lat"`
	Lng     int    `json:"lng"`
}
