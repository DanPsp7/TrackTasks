package models

type People struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Address        string `json:"address"`
	PassportNumber int    `json:"passport_number"`
}
