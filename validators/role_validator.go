package validators

import (
	"user-rest-app/models"
)

//var validate = validator.New()

func ValidateRole(role models.Role) error {
	return validate.Struct(role)
}
