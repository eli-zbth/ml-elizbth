package validations

import (
	"errors"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	englishTranslations "github.com/go-playground/validator/v10/translations/en"

)


type customValidator struct {
	validate   *validator.Validate
	translator ut.Translator
}

type CustomValidator interface {
	Validate(i interface{}) error
}


func NewCustomValidator(validate *validator.Validate) (CustomValidator, error) {
	cv := &customValidator{validate: validate}
	english := en.New()
	universalTranslator := ut.New(english, english)
	englishTranslator, _ := universalTranslator.GetTranslator("en")
	_ = englishTranslations.RegisterDefaultTranslations(validate, englishTranslator)
	cv.translator = englishTranslator

	
	return cv, nil
}

func (cv *customValidator) Validate(i interface{}) error {
	if i == nil {
		return errors.New("request has empty body")
	}
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.ValueOf(i).IsNil() {
			return errors.New("request has empty body")
		}
	}
	err := cv.validate.Struct(i)
	if err != nil {
		for _, errValidation := range err.(validator.ValidationErrors) {
			message := errValidation.Translate(cv.translator)
			return errors.New(message)
          
		}
        return err
	}
	return nil
}

