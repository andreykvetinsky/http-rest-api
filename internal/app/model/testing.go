package model

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestUser(t *testing.T) *User {

	return &User{
		Email:    "testi" + strconv.Itoa(rand.Int()) + "@gmail.com",
		Password: "password",
	}
}

func TestNote(t *testing.T) *Note {

	return &Note{
		User_ID: 1779,
		Note:    "conqquer Everest",
	}
}
