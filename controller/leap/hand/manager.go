package hand

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// Manager handle the hand triggered by leap motion device
type Manager interface {
	// ProcessGestures handle the gesture triggered by leap motion device
	ProcessHands(h leap.Hand)
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
	forward := NewForwardListener(l.c)
	backward := NewBackwardListener(l.c)
	forward.SetNext(backward)
	l.event = forward
}

func (l *leapHandManager) ProcessHands(hand leap.Hand) {
	l.event.Process(hand)
}

