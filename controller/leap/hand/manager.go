package hand

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// Manager handle the hand triggered by leap motion device
type Manager interface {
	// Process handle the gesture triggered by leap motion device
	Process(h leap.Hand)
}

type leapHandManager struct {
	c     *tello.Driver
	event EventListener
}

// NewManager listen the events related to leap device.
func NewManager(c *tello.Driver) Manager {
	listener := &leapHandManager{c: c}
	listener.build()
	return listener
}

func (l *leapHandManager) build() {
	/*forward := NewForwardListener(l.c)

	backward := NewBackwardListener(l.c)
	forward.SetNext(backward)
	*/
	right := NewRightListener(l.c)
	//	backward.SetNext(right)

	//left := NewLeftListener(l.c)
	//right.SetNext(left)

	l.event = right
}

func (l *leapHandManager) Process(hand leap.Hand) {
	l.event.Process(hand)
}
