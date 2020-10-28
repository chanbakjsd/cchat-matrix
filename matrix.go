package matrix

import (
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/services"
	"github.com/diamondburned/cchat/text"
	"github.com/diamondburned/cchat/utils/empty"

	"github.com/chanbakjsd/cchat-matrix/internal/auth"
)

func init() {
	services.RegisterService(Service{})
}

type Service struct {
	empty.Service
}

func (Service) Name() text.Rich {
	return text.Plain("Matrix")
}

func (Service) AsIconer() cchat.Iconer {
	return Iconer{}
}

func (Service) Authenticate() []cchat.Authenticator {
	return []cchat.Authenticator{auth.HomeServer{}}
}

func (Service) AsSessionRestorer() cchat.SessionRestorer {
	return SessionRestorer{}
}
