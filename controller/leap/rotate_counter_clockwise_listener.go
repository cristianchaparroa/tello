package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

type RotateCounterClockwiseListener struct {
	next EventListener
	c    *Controller
}

// NewRotateCounterClockwiseListener generates a pointer to RotateCounterClockwiseListener
func NewRotateCounterClockwiseListener(c *Controller) *RotateCounterClockwiseListener {
	return &RotateCounterClockwiseListener{c: c}
}

// Process verifies if is forward event and trigger it.
func (l *RotateCounterClockwiseListener) Process(gesture leap.Gesture) {

	if l.IsForwardEvent(gesture) {
		fmt.Println("IsRotateCounterClockwise")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(gesture)
}

// SetNext saves the next listener in the chain
func (l *RotateCounterClockwiseListener) SetNext(next EventListener) {
	l.next = next
}

// IsForwardEvent determins if is  land event.
func (l *RotateCounterClockwiseListener) IsForwardEvent(gesture leap.Gesture) bool {
	return false
}
