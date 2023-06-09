package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

func ValidateCreateTask(request srv.TaskCreateReq) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Detail, validation.Required),
	)
	if err != nil {
		panic(util.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateUpdateTask(request srv.TaskUpdateReq) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
	)
	if err != nil {
		panic(util.ValidationError{
			Message: err.Error(),
		})
	}
}
