package validation

import (
	"encoding/json"
	"errors"
	resterr "news-api/configuration/rest_err"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transl, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *resterr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError("invalid JSON provided")
	} else if errors.As(validation_err, &jsonValidationErr) {
		var causes []resterr.Cause
		for _, err := range jsonValidationErr {
			causes = append(causes, resterr.Cause{
				Field:   err.Field(),
				Message: err.Translate(transl),
			})
		}
		return resterr.NewBadRequestValidationError("validation error", causes)
	} else {
		return resterr.NewBadRequestError("internal server error")
	}
}
