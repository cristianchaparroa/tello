package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

type RotateClockwiseListener struct {
	next EventListener
	c    *Controller
}

// NewRotateClockwiseListener generates a pointer to RotateClockwiseListener
func NewRotateClockwiseListener(c *Controller) *RotateClockwiseListener {
	return &RotateClockwiseListener{c: c}
}

// Process verifies if is forward event and trigger it.
func (l *RotateClockwiseListener) Process(gesture leap.Gesture) {

	if l.IsForwardEvent(gesture) {
		fmt.Println("IsRotateClockwise")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(gesture)
}

// SetNext saves the next listener in the chain
func (l *RotateClockwiseListener) SetNext(next EventListener) {
	l.next = next
}

// IsForwardEvent determins if is  land event.
func (l *RotateClockwiseListener) IsForwardEvent(gesture leap.Gesture) bool {
	return false
}
