package models

type User struct {
	Id           uint
	Firstname    string
	Lastname     string
	Email        string
	Password     []byte
	IsAmbassador bool
}
