package controller

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

const (
	// DefaultShift is the shift value in case that there is not a configuration.
	DefaultShift = 20
)

// KeyConfig contains the basic data to setup the command in the drone
type KeyConfig struct {
	// Shift is the unit of desplacement used in different commands.
	Shift int
}

// KeyBoard is in charge to handle the keyboards commands
// and interact with drone
type KeyBoard struct {
	drone *tello.Driver
	keys  *keyboard.Driver
}

// NewKeyBoard generates a pointer to NewKeyBoard
func NewKeyBoard(drone *tello.Driver, keys *keyboard.Driver) *KeyBoard {
	fmt.Println(drone.Name())
	return &KeyBoard{drone: drone, keys: keys}
}

// Run function create the infinity loop to read all keys until the user
// wants to land the drone
func (c *KeyBoard) Run() {

	c.keys.On(keyboard.Key, func(data interface{}) {
		key := data.(keyboard.KeyEvent)
		switch key.Key {
		case keyboard.A:
			c.TurnLeft()
		case keyboard.D:
			c.TurnRight()
		case keyboard.W:
			c.Up()
		case keyboard.S:
			c.Down()
		case keyboard.Q:
			c.drone.Land()
		case keyboard.P:
			c.drone.TakeOff()
		case keyboard.ArrowUp:
			c.Forward()
		case keyboard.ArrowDown:
			fmt.Println(key.Char)
			c.Backward()
		case keyboard.ArrowLeft:
			c.Left()
		case keyboard.ArrowRight:
			c.Right()
		case keyboard.Escape:
			resetDronePostion(c.drone)
		}
	})
}

func resetDronePostion(drone *tello.Driver) {
	drone.Forward(0)
	drone.Backward(0)
	drone.Up(0)
	drone.Down(0)
	drone.Left(0)
	drone.Right(0)
	drone.Clockwise(0)
}

// Up moves up the drone
func (c *KeyBoard) Up() {
	c.drone.Up(DefaultShift)
}

// Down moves down the drone
func (c *KeyBoard) Down() {
	c.drone.Down(DefaultShift)
}

// TurnRight rotates the drone to the right side.
func (c *KeyBoard) TurnRight() {
	c.drone.Clockwise(DefaultShift)
}

// TurnLeft rotates the drone to the left side
func (c *KeyBoard) TurnLeft() {
	c.drone.Clockwise(-DefaultShift)
}

// Right moves right the drone
func (c *KeyBoard) Right() {
	c.drone.Right(DefaultShift)
}

// Left moves to left the drone
func (c *KeyBoard) Left() {
	c.drone.Left(DefaultShift)
}

// Forward drone
func (c *KeyBoard) Forward() {
	c.drone.Forward(DefaultShift)
}

// Backward drone
func (c *KeyBoard) Backward() {
	c.drone.Backward(DefaultShift)
}

// TakeOff start the drone and go up it
func (c *KeyBoard) TakeOff() {
	c.drone.TakeOff()
}

// Land the drone
func (c *KeyBoard) Land() {
	gobot.After(5*time.Second, func() {
		c.drone.Land()
	})
}
