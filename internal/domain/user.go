package domain

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	CardPin int    `json:"card_pin"`
}
