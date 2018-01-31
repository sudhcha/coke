package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/sudhcha/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
