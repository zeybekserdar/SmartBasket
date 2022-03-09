package product

import (
	turkish "github.com/go-playground/locales/tr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	trvalidator "github.com/go-playground/validator/v10/translations/tr"
)

type ErrorResponse struct {
	ErrorMessageTr string
	ErrorMessageEn string
}

func ValidateProduct(product Product) []*ErrorResponse {
	var errors []*ErrorResponse

	uni := ut.New(turkish.New())
	trans, _ := uni.GetTranslator("tr")
	validate := validator.New()
	trvalidator.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(product)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.ErrorMessageEn = err.Error()
			element.ErrorMessageTr = err.Translate(trans)
			errors = append(errors, &element)

		}
	}
	return errors
}
