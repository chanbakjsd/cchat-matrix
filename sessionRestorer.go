package matrix

import (
	"github.com/chanbakjsd/gotrix"
	"github.com/chanbakjsd/gotrix/matrix"
	"github.com/diamondburned/cchat"

	"github.com/chanbakjsd/cchat-matrix/internal/session"
)

type SessionRestorer struct{}

func (SessionRestorer) RestoreSession(storage map[string]string) (cchat.Session, error) {
	cli, err := gotrix.New(storage["homeserver"])
	// TODO: Fix this when there's a better API.
	cli.UserID = matrix.UserID(storage["userID"])
	cli.AccessToken = storage["accessToken"]

	// Do a quick sanity check.
	_, err = cli.Whoami()
	if err != nil {
		// This basically should never error if token is valid.
		return nil, err
	}

	return session.New(cli)
}
