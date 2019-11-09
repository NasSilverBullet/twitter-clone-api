package frameworks

import (
	"github.com/NasSilverBullet/twitter-clone-api/app/interfaces"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	*validator.Validate
}

func NewValidator() interfaces.Validator {
	return &Validator{validator.New()}
}

func (v *Validator) Struct(s interface{}) error {
	if err := v.Validate.Struct(s); err != nil {
		return err
	}

	return nil
}
