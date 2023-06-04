package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

func ValidateAuthRegister(request srv.AuthRegister) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.UserName, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
	if err != nil {
		panic(util.ValidationError{
			Message: err.Error(),
		})
	}
}
