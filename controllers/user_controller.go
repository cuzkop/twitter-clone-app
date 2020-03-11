package controllers

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func NewUserController() User {
	return User{}
}

func (u *User) Create(name string) *User {
	u.FirstName = name
	u.LastName = "takahashi"
	u.ID = 829
	return u
}
