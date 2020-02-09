package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

type ForwardListener struct {
	next EventListener
	c    *Controller
}

// NewForwardListener generates a pointer to ForwardListener
func NewForwardListener(c *Controller) *ForwardListener {
	return &ForwardListener{c: c}
}

// Process verifies if is forward event and trigger it.
func (l *ForwardListener) Process(gesture leap.Gesture) {

	if l.IsForwardEvent(gesture) {
		fmt.Println("IsForward")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(gesture)
}

// SetNext saves the next listener in the chain
func (l *ForwardListener) SetNext(next EventListener) {
	l.next = next
}

// IsForwardEvent determins if is  land event.
func (l *ForwardListener) IsForwardEvent(gesture leap.Gesture) bool {
	return false
}
