package auth

import (
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/text"
	"maunium.net/go/mautrix"
)

type TokenAuth struct {
	Client        *mautrix.Client
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
		cchat.AuthenticateEntry{
			Name:   "Token",
			Secret: true,
		},
	}
}

func (p TokenAuth) Authenticate(s []string) (cchat.Session, cchat.AuthenticateError) {
	token := s[0]

	return login(p.Client, &mautrix.ReqLogin{
		Type:                     mautrix.AuthTypeToken,
		Token:                    token,
		InitialDeviceDisplayName: "cchat-matrix",

		StoreCredentials: true,
	})
}
