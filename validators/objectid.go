package validators

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterObjectIDValidator(v *validator.Validate) error {
	if err := v.RegisterValidation("objectid", isObjectID); err != nil {
		return err
	}

	return nil
}

func isObjectID(fl validator.FieldLevel) bool {
	return primitive.IsValidObjectID(fl.Field().String())
}
