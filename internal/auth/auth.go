package auth

import (
	"github.com/chanbakjsd/gotrix/api"
	"github.com/diamondburned/cchat"
	"github.com/pkg/errors"
)

func processLoginErrors(err error) cchat.AuthenticateError {
	if err == nil {
		return nil
	}

	switch err {
	case api.ErrInvalidRequest:
		return cchat.WrapAuthenticateError(err)
	case api.ErrInvalidCreds:
		return cchat.WrapAuthenticateError(errors.New("invalid credentials provided"))
	case api.ErrUserDeactivated:
		return cchat.WrapAuthenticateError(errors.New("user has been deactivated"))
	default:
		return cchat.WrapAuthenticateError(err)
	}
}
