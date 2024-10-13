package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

var (
	validate         *validator.Validate
	errorMessagesMap = map[string]string{
		"required": "必須です",
		"email":    "emailの形にしてください",
	}
)

func Init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func GenerateValidationErrorMessages(err error) ErrorResponse {
	errorMessages := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		message := getErrorMessage(err)
		errorMessages[strings.ToLower(err.Field())] = message
	}
	return ErrorResponse{Errors: errorMessages}
}
func getErrorMessage(err validator.FieldError) string {
	if msg, exists := errorMessagesMap[err.Tag()]; exists {
		return msg
	}
	return "無効な値です"
}
