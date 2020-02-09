package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

type BackwardListener struct {
	next EventListener
	c    *Controller
}

// NewBackwardListener generates a pointer to BackwardListener
func NewBackwardListener(c *Controller) *BackwardListener {
	return &BackwardListener{c: c}
}

// Process verifies if is backward event and trigger it.
func (l *BackwardListener) Process(gesture leap.Gesture) {

	if l.IsBackwardEvent(gesture) {
		fmt.Println("IsLand")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(gesture)
}

// SetNext saves the next listener in the chain
func (l *BackwardListener) SetNext(next EventListener) {
	l.next = next
}

func (l *BackwardListener) IsBackwardEvent(gesture leap.Gesture) bool {
	return false
}
