package app

import "example.com/testproject/domain"

func WithZeroValue() {
	var _ domain.User
	var _ *domain.User
}

func WithZeroValueAndConstructor(isMale bool) {
	var user *domain.User
	if isMale {
		user, _ = domain.NewUser(1, "Alice", 30)
	} else {
		user, _ = domain.NewUser(2, "Bob", 25)
	}
	_ = user
}

type alias = domain.User

func WithTypeAlias() {
	_ = alias{
		ID:   123,
		Name: "Eve",
		Age:  28,
	}
}
