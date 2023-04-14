package models

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
	City    string `json:"city"`
}
