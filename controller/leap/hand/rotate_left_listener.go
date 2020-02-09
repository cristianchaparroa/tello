package hand

import (
	"fmt"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

type RotateLeftListener struct {
	next EventListener
	c    *tello.Driver
}

// NewTurnLeftListener generates a pointer to RotateLeftListener
func NewTurnLeftListener(c *tello.Driver) *RotateLeftListener {
	return &RotateLeftListener{c: c}
}

// Process verifies if is forward event and trigger it.
func (l *RotateLeftListener) Process(hand leap.Hand) {

	if l.IsRotateLeftEvent(hand) {
		fmt.Println("IsRotateCounterClockwise")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(hand)
}

// SetNext saves the next listener in the chain
func (l *RotateLeftListener) SetNext(next EventListener) {
	l.next = next
}

// IsRotateRight determines if is  land event.
func (l *RotateLeftListener) IsRotateLeftEvent(hand leap.Hand) bool {
	return false
}
