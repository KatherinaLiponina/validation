package validation

import (
	"errors"
	"reflect"
)

var ErrFailedValidation = errors.New("validation of struct have failed")

func Validate(v any) error {
	if reflect.ValueOf(v).Kind() != reflect.Struct {
		return ErrFailedValidation
	}

	dt := reflect.TypeOf(v)
	dv := reflect.ValueOf(v)

	for i := 0; i < dt.NumField(); i++ {

		validationGoal := dt.Field(i).Tag.Get("validate")
		if validationGoal == "" {
			continue
		}

		if !dt.Field(i).IsExported() {
			return ErrFailedValidation
		}

		if dt.Field(i).Type.Kind() == reflect.String {
			err := validateString(validationGoal, dv.Field(i).Interface().(string))
			if err != nil {
				return ErrFailedValidation
			}
		}
	}

	return nil
}

func validateString(validation string, value string) error {
	switch validation {
	case "title":
		if len(value) <= 0 || len(value) >= 100 {
			return ErrFailedValidation
		}
	case "text":
		if len(value) <= 0 || len(value) >= 500 {
			return ErrFailedValidation
		}
	}
	return nil
}