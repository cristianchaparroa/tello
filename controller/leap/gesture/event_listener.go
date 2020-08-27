package gesture

import "gobot.io/x/gobot/platforms/leap"

// EventListener set the methods to process
// different events.
type EventListener interface {
	Process(gesture leap.Gesture) bool
	SetNext(l EventListener)
}
