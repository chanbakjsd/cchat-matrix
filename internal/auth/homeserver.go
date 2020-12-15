package auth

import (
	"github.com/chanbakjsd/gotrix"
	"github.com/chanbakjsd/gotrix/matrix"
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/text"
	"github.com/pkg/errors"
)

type HomeServer struct{}

func (HomeServer) Name() text.Rich {
	return text.Plain("Homeserver")
}

func (HomeServer) Description() text.Rich {
	return text.Plain("Enter the address of your homeserver.")
}

func (HomeServer) AuthenticateForm() []cchat.AuthenticateEntry {
	return []cchat.AuthenticateEntry{
		{
			Name:        "Address",
			Placeholder: "matrix.org",
			Description: "The domain or IP associated with your homeserver.",
		},
	}
}

func (HomeServer) Authenticate(s []string) (cchat.Session, cchat.AuthenticateError) {
	client, err := gotrix.New(s[0])
	if err != nil {
		return nil, cchat.WrapAuthenticateError(errors.Wrap(err, "cannot connect to homeserver"))
	}
	methods, err := client.GetLoginMethods()
	if err != nil {
		return nil, cchat.WrapAuthenticateError(errors.Wrap(err, "cannot fetch login methods"))
	}

	supportedMethods := make([]cchat.Authenticator, 0, 2)
	for _, v := range methods {
		switch v {
		case matrix.LoginPassword:
			supportedMethods = append(supportedMethods, PasswordAuth{
				Client: client,
			})
		case matrix.LoginToken:
			supportedMethods = append(supportedMethods, TokenAuth{
				Client: client,
			})
		}
	}

	if len(supportedMethods) == 0 {
		return nil, cchat.WrapAuthenticateError(errors.Wrap(err, "no supported login methods found"))
	}

	return nil, NextStage{
		AuthMethods: supportedMethods,
		Message:     "Please login using your preferred method.",
	}
}
