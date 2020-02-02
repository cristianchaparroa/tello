package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

// Logger is an utility to show information related to different leap motion
// structures
type Logger interface {
	ShowGesture(g leap.Gesture)
}

type logger struct {
}

// NewLogger generates a pointer to logger.
func NewLogger() Logger {
	return &logger{}
}

// ShowGesture show the information related to gesture
func (l *logger) ShowGesture(g leap.Gesture) {
	message := fmt.Sprintf("Type:%s, State:%s, IsInProggress:%v, Normal:%v", g.Type, g.State, g.Progress, g.Normal)
	fmt.Println(message)
}
