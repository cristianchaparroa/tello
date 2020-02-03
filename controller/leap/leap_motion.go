package leap

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// Controller is a Controller in charge to handle the
// moviements of tello drone with Controller device.
type Controller struct {
	drone  *tello.Driver
	leap   *leap.Driver
	logger Logger
}

// NewController generates pointer to Controller
func NewController(drone *tello.Driver, leap *leap.Driver) *Controller {
	logger := NewLogger()
	c := &Controller{drone: drone, leap: leap, logger: logger}
	return c
}

// Run function create the infinity loop to read all Frames until the user
// wants to land the drone
func (c *Controller) Run() {

	c.leap.On(leap.GestureEvent, func(data interface{}) {

		gesture := data.(leap.Gesture)
		c.logger.ShowGesture(gesture)
		listener := NewleapMotionListener(c)
		listener.ProcessGestures(gesture)
	})

}

// TakeOff the drone with the
func (c *Controller) TakeOff() {
	c.drone.TakeOff()
}

// Up  shift the drone up
func (c *Controller) Up() {

}

// Down  shift the drone down
func (c *Controller) Down() {}

// Right shift the drone right
func (c *Controller) Right() {}

// Left  shift the drone  left.
func (c *Controller) Left() {}

// Land the drone
func (c *Controller) Land() {}
