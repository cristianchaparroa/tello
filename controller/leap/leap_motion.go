package leap

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

const (
	// DefaultShift is the shift value in case that there is not a configuration.
	DefaultShift = 20
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
		manager := NewleapGestureManager(c)
		manager.ProcessGestures(gesture)
	})

}

// TakeOff the drone with the
func (c *Controller) TakeOff() {
	c.drone.TakeOff()
}

// Up  shift the drone up
func (c *Controller) Up() {
	c.drone.Up(DefaultShift)
}

// Down  shift the drone down
func (c *Controller) Down() {
	c.drone.Down(DefaultShift)
}

// Right shift the drone right
func (c *Controller) Right() {
	c.drone.Right(DefaultShift)
}

// Left  shift the drone  left.
func (c *Controller) Left() {
	c.drone.Left(DefaultShift)
}

// Land the drone
func (c *Controller) Land() {
	gobot.After(5*time.Second, func() {
		c.drone.Land()
	})
}
