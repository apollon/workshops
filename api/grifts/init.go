package grifts

import (
	"github.com/apollon/workshops/api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
