package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

type RightListener struct {
	next EventListener
	c    *Controller
}

func NewRightListener(c *Controller) *RightListener {
	return &RightListener{c: c}
}

func (l *RightListener) Process(gesture leap.Gesture) {

	if l.IsRightEvent(gesture) {
		fmt.Println("IsRightMovement")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(gesture)
}

func (l *RightListener) IsRightEvent(gesture leap.Gesture) bool {
	return false
}

func (l *RightListener) SetNext(next EventListener) {
	l.next = next
}
