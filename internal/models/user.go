package models

type User struct {
	Model
	Name string
	Age  int
}

func NewUser(name string, age int) *User {

	user := &User{
		Name: name,
		Age:  age,
	}
	user.Init("users")

	return user
}
