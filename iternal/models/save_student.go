package models

type SaveStudent struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Country string `json:"country"`
	Email   string `json:"email"`
}
