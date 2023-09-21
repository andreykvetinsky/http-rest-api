package model

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestUser(t *testing.T) *User {
	random := rand.Intn(100)
	s := "test" + strconv.Itoa(random) + "@example.com"
	return &User{
		Email:    s,
		Password: "password",
	}
}
