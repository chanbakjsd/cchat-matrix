package session

import (
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/text"
	"github.com/diamondburned/cchat/utils/empty"
	"maunium.net/go/mautrix"
)

type Session struct {
	empty.Session
	*mautrix.Client
	ShouldStop bool
}

func New(cli *mautrix.Client) cchat.Session {
	s := &Session{
		Client: cli,
	}
	go s.Listen()
	return s
}

func (s *Session) ID() cchat.ID {
	return cchat.ID(s.UserID)
}

func (s *Session) Name() text.Rich {
	return text.Plain(string(s.UserID))
}

func (s *Session) AsIconer() cchat.Iconer {
	//TODO stub
	return nil
}

func (s *Session) Servers(cchat.ServersContainer) error {
	//TODO stub
	return nil
}

func (s *Session) Disconnect() error {
	s.ShouldStop = true
	s.StopSync()
	return nil
}

func (s *Session) AsCommander() cchat.Commander {
	//TODO stub
	return nil
}

func (s *Session) AsSessionSaver() cchat.SessionSaver {
	return s
}

func (s *Session) SaveSession() map[string]string {
	return map[string]string{
		"homeserver":  s.HomeserverURL.String(),
		"accessToken": s.AccessToken,
		"deviceID":    string(s.DeviceID),
		"userID":      string(s.UserID),
	}
}
