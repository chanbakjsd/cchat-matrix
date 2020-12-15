package auth

import (
	"github.com/chanbakjsd/gotrix"
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/text"

	"github.com/chanbakjsd/cchat-matrix/internal/session"
)

type TokenAuth struct {
	Client        *gotrix.Client
	TransactionID string
}

func (TokenAuth) Name() text.Rich {
	return text.Plain("Login Token")
}

func (TokenAuth) Description() text.Rich {
	return text.Plain("Login using your token.")
}

func (TokenAuth) AuthenticateForm() []cchat.AuthenticateEntry {
	return []cchat.AuthenticateEntry{
		{
			Name:   "Token",
			Secret: true,
		},
	}
}

func (p TokenAuth) Authenticate(s []string) (cchat.Session, cchat.AuthenticateError) {
	token := s[0]

	err := p.Client.LoginToken(token)
	if err != nil {
		return nil, processLoginErrors(err)
	}

	sess, err := session.New(p.Client)
	if err != nil {
		return nil, cchat.WrapAuthenticateError(err)
	}
	return sess, nil
}
