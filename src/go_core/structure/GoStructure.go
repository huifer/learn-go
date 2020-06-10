package structure

// 用户结构
type User struct {
	Name string
	Age  int
	Pro  int
}

// 学生结构
type Student struct {
	// 指向用户
	User *User
	// 班级号
	ClassNum int
	Pro      int
}

func GetUser() *User {

	u := User{Age: 1, Name: "huifer"}
	return &u
}

func GetStudent() *Student {
	user := GetUser()
	s := Student{User: user, ClassNum: 10}
	return &s
}
