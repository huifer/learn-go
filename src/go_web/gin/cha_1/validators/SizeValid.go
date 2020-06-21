package validators

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var SizeValid validator.Func = func(fl validator.FieldLevel) bool {
	data := fl.Field().Interface().(int)
	if data > 100 {
		return false
	}
	return true

}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}
