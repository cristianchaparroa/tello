package leap

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

type leapMotionListener struct {
}

// NewleapMotionListener listen the events related to leap device.
func NewleapMotionListener() MotionListener {
	return &leapMotionListener{}
}

func (l *leapMotionListener) ProcessGestures(g leap.Gesture) {

}

func (l *leapMotionListener) IsTakeOffEvent(gesture leap.Gesture) bool {

	if GestureCircle != gesture.Type {
		return false
	}

	isClockWise := isClockWise(gesture)
	isTwoRounds := isTwoRounds(gesture)

	return isClockWise && isTwoRounds
}
