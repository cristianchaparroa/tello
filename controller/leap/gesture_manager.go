package leap

import "gobot.io/x/gobot/platforms/leap"

// GestureManager handle the gesture trigered by leap motion device
type GestureManager interface {

	// ProcessGestures handle the gestures trigered by leap motion device
	ProcessGestures(g leap.Gesture)
}
