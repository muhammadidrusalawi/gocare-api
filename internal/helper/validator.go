package helper

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s any) string {
	err := validate.Struct(s)
	if err == nil {
		return ""
	}

	var messages []string

	for _, e := range err.(validator.ValidationErrors) {
		messages = append(messages, e.Field()+" "+e.Tag())
	}

	return strings.Join(messages, ", ")
}
