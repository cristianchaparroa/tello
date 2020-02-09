package gesture

import "gobot.io/x/gobot/platforms/leap"

const (
	// CircleThreshold is the threshold to determine that the gesture is two rounds
	CircleThreshold = 1.5
)

// EventListener set the methods to process
// different events.
type EventListener interface {
	Process(gesture leap.Gesture)
	SetNext(l EventListener)
}
