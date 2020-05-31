package gesture

import (
	"fmt"

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

// isClockWise verifies if gesture is clockwise.
func isClockWise(gesture leap.Gesture) bool {
	normal := gesture.Normal
	isNormal := len(normal) == 3

	clockwise := false
	if isNormal && normal[ZAxis] < 0 {
		clockwise = true
	}

	return clockwise
}

// isTwoRounds determines if the gesture is two rounds
func isTwoRounds(gesture leap.Gesture) bool {
	fmt.Printf("Progress:%v\n", gesture.Progress)
	return gesture.Progress > CircleThreshold
}
