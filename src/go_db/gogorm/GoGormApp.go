package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Udemo struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("mysql", "huifer:a12345@tcp(47.98.225.144:3306)/scrum?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	//
	//db.CreateTable(&Udemo{})

	var d = Udemo{
		Model: gorm.Model{},
		Name:  "zhangsan",
	}
	db.Create(&d)
	db.Delete(&d)

}
