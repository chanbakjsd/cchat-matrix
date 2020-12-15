package auth

import (
	"github.com/chanbakjsd/gotrix"
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/text"

	"github.com/chanbakjsd/cchat-matrix/internal/session"
)

type PasswordAuth struct {
	Client *gotrix.Client
}

func (PasswordAuth) Name() text.Rich {
	return text.Plain("Password")
}

func (PasswordAuth) Description() text.Rich {
	return text.Plain("Login using your password.")
}

func (PasswordAuth) AuthenticateForm() []cchat.AuthenticateEntry {
	return []cchat.AuthenticateEntry{
		{
			Name:        "User ID",
			Description: "The user ID you entered when you registered on your Matrix homeserver.",
		},
		{
			Name:   "Password",
			Secret: true,
		},
	}
}

func (p PasswordAuth) Authenticate(s []string) (cchat.Session, cchat.AuthenticateError) {
	username := s[0]
	password := s[1]

	err := p.Client.LoginPassword(username, password)
	if err != nil {
		return nil, processLoginErrors(err)
	}

	sess, err := session.New(p.Client)
	if err != nil {
		return nil, cchat.WrapAuthenticateError(err)
	}
	return sess, nil
}
