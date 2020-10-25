package appError

import (
	"errors"
	"fmt"
)

const (
	UnauthorizedMessage        = "user is not authorized to perform this operation"
	InternalServerErrorMessage = "sorry something went wrong"
	NotFoundMessage            = "not found"
	InvalidParameterMessage    = "is invalid"
	internal                   = "INTERNAL_ERROR"
	api                        = "API_ERROR"
)

func NotFound(prop string) error {
	return errors.New(fmt.Sprintf("%v %v", prop, NotFoundMessage))
}

func InvalidParameter(prop string) error {
	return errors.New(fmt.Sprintf("%v %v", prop, InvalidParameterMessage))
}

func Unauthorized() error {
	return errors.New(UnauthorizedMessage)
}

func InternalServer() error {
	return errors.New(InternalServerErrorMessage)
}
