package validation

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)
 
func IranianMobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	res, err := regexp.MatchString(`^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res

}
