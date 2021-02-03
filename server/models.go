package server

import (
	"fmt"
)

type Users struct {
	Users []User `json:"data"`
}

type User struct {
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Postcode  string `json:"postcode"`
	CreatedAt string `json:"created_at"`
}

type RandomUsers struct {
	RandomUsers []RandomUser `json:"results"`
}

type RandomUser struct {
	Gender string `json:"gender"`
	Name   struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Location     Location `json:"location"`
	Registration struct {
		Date string `json:"date"`
	} `json:"registered"`
}

type Location struct {
	Postcode interface{} `json:"postcode"`
}

func (l Location) getString() (string, error) {
	switch v := l.Postcode.(type) {
	case float64:
		return fmt.Sprintf("%f", v), nil
	case string:
		return v, nil
	}
	return "", fmt.Errorf("postcode isn't one of the types above")
}
