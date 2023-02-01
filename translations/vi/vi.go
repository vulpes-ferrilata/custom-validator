package vi

import (
	"log"

	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) error {
	t := struct {
		tag         string
		translation string
		override    bool
	}{
		tag:         "objectid",
		translation: "{0} phải là ObjectID hợp lệ",
		override:    false,
	}

	return v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		if err := ut.Add(tag, translation, override); err != nil {
			return err
		}

		return nil
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
