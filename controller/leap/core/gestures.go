package core

import (
	"gobot.io/x/gobot/platforms/leap"
)

const (
	// CircleThreshold is the threshold to determine that the gesture is two rounds
	CircleThreshold = 1
)

// IsCircleGesture determines if the gesture is a circle
func IsCircleGesture(gesture leap.Gesture) bool {
	return Circle == gesture.Type
}

// IsClockWise verifies if gesture is clockwise.
func IsClockWise(gesture leap.Gesture) bool {
	normal := gesture.Normal
	isNormal := len(normal) == 3

	clockwise := false
	if isNormal && normal[ZAxis] < 0 {
		clockwise = true
	}

	return clockwise
}

// IsTwoRounds determines if the gesture is two rounds
func IsTwoRounds(gesture leap.Gesture) bool {
	return gesture.Progress > CircleThreshold
}
