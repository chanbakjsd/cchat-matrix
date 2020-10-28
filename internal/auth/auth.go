package auth

import (
	"github.com/diamondburned/cchat"
	"github.com/pkg/errors"
	"maunium.net/go/mautrix"

	"github.com/chanbakjsd/cchat-matrix/internal/session"
)

func login(client *mautrix.Client, req *mautrix.ReqLogin) (cchat.Session, cchat.AuthenticateError) {
	_, err := client.Login(req)

	if err != nil {
		respError, ok := err.(mautrix.RespError)
		if !ok {
			return nil, cchat.WrapAuthenticateError(errors.Wrap(err, "error logging in"))
		}

		switch {
		case respError.Is(mautrix.MForbidden):
			return nil, cchat.WrapAuthenticateError(errors.New("invalid credentials provided"))
		case respError.Is(mautrix.MUserDeactivated):
			return nil, cchat.WrapAuthenticateError(errors.New("user has been deactivated"))
		case respError.Is(mautrix.MLimitExceeded):
			return nil, cchat.WrapAuthenticateError(errors.New("rate limit has been exceeded. try again later"))
		}

		return nil, cchat.WrapAuthenticateError(respError)
	}

	return session.New(client), nil
}
