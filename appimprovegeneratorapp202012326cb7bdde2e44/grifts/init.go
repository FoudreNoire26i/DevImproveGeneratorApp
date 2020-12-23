package grifts

import (
	"generatedApps/appimprovegeneratorapp202012326cb7bdde2e44/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
