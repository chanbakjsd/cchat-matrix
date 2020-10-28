package auth

import (
	"github.com/diamondburned/cchat"
)

type NextStage struct {
	AuthMethods []cchat.Authenticator
	Message     string
}

func (n NextStage) NextStage() []cchat.Authenticator {
	return n.AuthMethods
}

func (n NextStage) Error() string {
	return n.Message
}
