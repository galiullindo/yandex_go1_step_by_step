package main

import "errors"

var (
	ErrNameIsEmpty = errors.New("name is empty for user")
)

type User struct {
	Name     string
	Age      int
	IsActive bool
}

func NewUser(name string, age int) (*User, error) {
	if name == "" {
		return nil, ErrNameIsEmpty
	}
	if age == 0 {
		age = 18
	}
	return &User{Name: name, Age: age, IsActive: true}, nil
}
