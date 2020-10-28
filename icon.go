package matrix

import (
	"context"

	"github.com/diamondburned/cchat"
)

const (
	IconSource = "https://raw.githubusercontent.com/vector-im/logos/master/matrix/matrix-logo.png"
)

type Iconer struct{}

func (Iconer) Icon(_ context.Context, ic cchat.IconContainer) (stop func(), err error) {
	ic.SetIcon(IconSource)
	return func() {}, nil
}
