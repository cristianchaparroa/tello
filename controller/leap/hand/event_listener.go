package hand

import "gobot.io/x/gobot/platforms/leap"

const (
	TurnThreshold        = 0.2
	TurnSpeedFactor      = 2.0
	DirectionThreshold   = 0.25
	DirectionSpeedFactor = 0.05
	UpControlThreshold   = 50
	UpSpeedFactor        = 0.01
	// CircleThreshold is the threshold to determine that the gesture is two rounds
	CircleThreshold = 1.5
)

// EventListener set the methods to process
// different events.
type EventListener interface {
	Process(hand leap.Hand)
	SetNext(l EventListener)
}
