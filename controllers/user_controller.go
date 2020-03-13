package controllers

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
