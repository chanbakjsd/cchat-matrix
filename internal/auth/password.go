package auth

import (
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/text"
	"maunium.net/go/mautrix"
)

type PasswordAuth struct {
	Client *mautrix.Client
}

func (PasswordAuth) Name() text.Rich {
	return text.Plain("Password")
}

func (PasswordAuth) Description() text.Rich {
	return text.Plain("Login using your password.")
}

func (PasswordAuth) AuthenticateForm() []cchat.AuthenticateEntry {
	return []cchat.AuthenticateEntry{
		cchat.AuthenticateEntry{
			Name:        "User ID",
			Description: "The user ID you entered when you registered on your Matrix homeserver.",
		},
		cchat.AuthenticateEntry{
			Name:   "Password",
			Secret: true,
		},
	}
}

func (p PasswordAuth) Authenticate(s []string) (cchat.Session, cchat.AuthenticateError) {
	username := s[0]
	password := s[1]

	return login(p.Client, &mautrix.ReqLogin{
		Type: mautrix.AuthTypePassword,
		Identifier: mautrix.UserIdentifier{
			Type: mautrix.IdentifierTypeUser,
			User: username,
		},
		Password:                 password,
		InitialDeviceDisplayName: "cchat-matrix",

		StoreCredentials: true,
	})
}
