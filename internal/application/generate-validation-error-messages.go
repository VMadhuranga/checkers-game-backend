package application

import "github.com/go-playground/validator/v10"

func appendValidationErrorMessage(field []string, err validator.FieldError) []string {
	return append(field, validationErrorMessages[validationError{err.StructField(), err.ActualTag()}])
}

func generateValidationErrorMessages(errors validator.ValidationErrors) validationErrorMessagesResponse {
	messages := validationErrorMessagesResponse{}

	for _, err := range errors {
		switch err.StructField() {
		case "Username":
			messages.Username = appendValidationErrorMessage(messages.Username, err)
		case "Password":
			messages.Password = appendValidationErrorMessage(messages.Password, err)
		case "ConfirmPassword":
			messages.ConfirmPassword = appendValidationErrorMessage(messages.ConfirmPassword, err)
		case "NewUsername":
			messages.NewUsername = appendValidationErrorMessage(messages.NewUsername, err)
		}
	}

	return messages
}
