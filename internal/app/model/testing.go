package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "test6@example.com",
		Password: "password",
	}
}
