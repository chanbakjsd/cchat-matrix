package matrix

import (
	"github.com/diamondburned/cchat"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"

	"github.com/chanbakjsd/cchat-matrix/internal/session"
)

type SessionRestorer struct{}

func (SessionRestorer) RestoreSession(storage map[string]string) (cchat.Session, error) {
	cli, err := mautrix.NewClient(storage["homeserver"], id.UserID(storage["userID"]), storage["accessToken"])
	if err != nil {
		return nil, err
	}

	// Do a quick sanity check.
	_, err = cli.Whoami()
	if err != nil {
		// This basically should never error if token is valid.
		return nil, err
	}

	return session.New(cli), nil
}
