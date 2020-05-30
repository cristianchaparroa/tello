package hand

import (
	"fmt"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

type RotateRightListener struct {
	next EventListener
	c    *tello.Driver
}

// NewRotateRightListener generates a pointer to RotateRightListener
func NewRotateRightListener(c *tello.Driver) *RotateRightListener {
	return &RotateRightListener{c: c}
}

// Process verifies if is forward event and trigger it.
func (l *RotateRightListener) Process(hand leap.Hand) {

	if l.IsRotateRight(hand) {
		fmt.Println("IsRotateRight")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(hand)
}

// SetNext saves the next listener in the chain
func (l *RotateRightListener) SetNext(next EventListener) {
	l.next = next
}

// IsRotateRight determines if is  rotate clockwise event.
func (l *RotateRightListener) IsRotateRight(hand leap.Hand) bool {
	return false
}
