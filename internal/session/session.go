package session

import (
	"github.com/chanbakjsd/gotrix"
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/text"
	"github.com/diamondburned/cchat/utils/empty"
)

type Session struct {
	empty.Session
	*gotrix.Client
	ShouldStop bool
}

func New(cli *gotrix.Client) (cchat.Session, error) {
	s := &Session{
		Client: cli,
	}
	err := s.Open()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Session) ID() cchat.ID {
	return cchat.ID(s.UserID)
}

func (s *Session) Name() text.Rich {
	return text.Plain(string(s.UserID))
}

func (s *Session) AsIconer() cchat.Iconer {
	// TODO stub
	return nil
}

func (s *Session) Servers(cchat.ServersContainer) error {
	// TODO stub
	return nil
}

func (s *Session) Disconnect() error {
	return s.Close()
}

func (s *Session) AsCommander() cchat.Commander {
	// TODO stub
	return nil
}

func (s *Session) AsSessionSaver() cchat.SessionSaver {
	return s
}

func (s *Session) SaveSession() map[string]string {
	return map[string]string{
		"homeserver":  s.HomeServer,
		"accessToken": s.AccessToken,
		"deviceID":    string(s.DeviceID),
		"userID":      string(s.UserID),
	}
}

func (s *Session) Listen() {
	s.Open()
}
