package domain

import "math/rand/v2"

func CreateUser(name string, age int) (*User, error) {
	return NewUser(rand.Int(), name, age)
}

func CreateUserDirect(name string, age int) *User {
	return &User{ // want "direct construction of struct User is prohibited, use constructor function"
		ID:   rand.Int(),
		Name: name,
		Age:  age,
	}
}
