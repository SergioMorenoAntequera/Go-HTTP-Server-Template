package models

type User struct {
	Model
	Name string
}

func NewUser(name string) *User {

	user := &User{
		Name: name,
	}
	user.Init("user")

	return user
}
