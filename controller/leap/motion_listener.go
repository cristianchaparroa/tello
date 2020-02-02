package leap

import "gobot.io/x/gobot/platforms/leap"

// MotionListener handle the events trigered by leap motion device
type MotionListener interface {

	// ProcessGestures handle the gestures trigered by leap motion device
	ProcessGestures(g leap.Gesture)

	// IsTakeOffEvent it determines if is taoke off
	IsTakeOffEvent(gesture leap.Gesture) bool
}
