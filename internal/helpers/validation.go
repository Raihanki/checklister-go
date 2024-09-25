package helpers

import "github.com/go-playground/validator/v10"

func ValidateStruct(data interface{}) []ValidationErrorResponse {
	var validate = validator.New()

	var validationErrors []ValidationErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Error:       err.Error(),
			})
		}
	}

	return validationErrors
}
