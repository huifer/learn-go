package main

import (
	_ "learn-go/src/pac/inpac"
	"learn-go/src/structure"

	_ "learn-go/src/structure"
)

func main() {
	//pac.Pac01()
	//inpac.Package02()
	//goStrings.UseStrings()

	user := structure.GetUser()
	print(user.Age)
	print(user.Name)

}
