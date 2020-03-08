package structure

type User struct {
	Name string
	Age  int
}

func GetUser() *User {
	u := User{Age: 1, Name: "huifer"}
	return &u
}
