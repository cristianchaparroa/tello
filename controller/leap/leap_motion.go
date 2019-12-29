package leap

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// LeapMotion is a controller in charge to handle the
// moviements of tello drone with LeapMotion device.
type LeapMotion struct {
	drone    *tello.Driver
	leap     *leap.Driver
	listener leapMotionListener
}

// NewLeapMotion generates pointer to LeapMotion
func NewLeapMotion(drone *tello.Driver, leap *leap.Driver) *LeapMotion {
	return &LeapMotion{drone: drone, leap: leap}
}

// Run function create the infinity loop to read all Frames until the user
// wants to land the drone
func (c *LeapMotion) Run() {

	c.leap.On(leap.GestureEvent, func(data interface{}) {
		gesture := data.(leap.Gesture)
		c.listener.ProcessGestures(gesture)
	})

}

// TakeOff the drone with the
func (c *LeapMotion) TakeOff() {
	c.drone.TakeOff()
}

// Up  shift the drone up
func (c *LeapMotion) Up() {}

// Down  shift the drone down
func (c *LeapMotion) Down() {}

// Right shift the drone right
func (c *LeapMotion) Right() {}

// Left  shift the drone  left.
func (c *LeapMotion) Left() {}

// Land the drone
func (c *LeapMotion) Land() {}
