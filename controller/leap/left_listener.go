package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

type LeftListener struct {
	next EventListener
	c    *Controller
}

func NewLeftListener(c *Controller) *LeftListener {
	return &LeftListener{c: c}
}

func (l *LeftListener) Process(gesture leap.Gesture) {

	if l.IsLeftEvent(gesture) {
		fmt.Println("IsLeftMovement")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(gesture)

}

func (l *LeftListener) SetNext(next EventListener) {
	l.next = next
}

func (l *LeftListener) IsLeftEvent(gesture leap.Gesture) bool {
	return false
}
