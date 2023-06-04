package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

func ValidateAuthRegister(request srv.AuthRegisterReq) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.RoleCode, validation.Required),
	)
	if err != nil {
		panic(util.ValidationError{
			Message: err.Error(),
		})
	}
}
