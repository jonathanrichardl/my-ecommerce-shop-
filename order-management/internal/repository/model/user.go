package model

import (
	"strings"
	"time"
)

type Role uint8

const (
	BUYER Role = iota
	SELLER
)

func GetRole(role string) Role {
	switch strings.ToLower(role) {
	case "buyer":
		return BUYER
	default:
		return SELLER
	}
}

type User struct {
	Id          string
	Username    string
	Password    string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	Role        Role
}
