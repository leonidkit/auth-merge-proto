package models

import "time"

type PersonalData struct {
	ID        int
	UserID    int
	StatusID  int
	Email     string
	FirtName  string
	LastName  string
	BirthDate time.Time
	Sex       UserGender
	Country   string
	Address   string
}

type UserGender struct {
	Male, Female, Undefined string
}
