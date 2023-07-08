package models

type Deposition struct {
	Id         int    `json:"id"`
	Picture    string `json:"picture"`
	Deposition string `json:"deposition"`
	Name       string `json:"name"`
}
