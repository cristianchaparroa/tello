package leap

import "gobot.io/x/gobot/platforms/leap"

// IsCircleGesture determines if the gesture is a circle
func IsCircleGesture(gesture leap.Gesture) bool {
	return GestureCircle == gesture.Type
}

// isClockWise verifies if gestures is clockwise.
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
	return gesture.Progress > CircleThreshold
}
