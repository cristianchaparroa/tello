package controller

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

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

	fmt.Println("<-- Run ")

	c.keys.On(keyboard.Key, func(data interface{}) {
		key := data.(keyboard.KeyEvent)
		switch key.Key {
		case keyboard.A:
			fmt.Println(key.Char)
			c.drone.Clockwise(-25)
		case keyboard.D:
			fmt.Println(key.Char)
			c.drone.Clockwise(25)
		case keyboard.W:
			fmt.Println(key.Char)
			c.drone.Up(20)
		case keyboard.S:
			fmt.Println(key.Char)
			c.drone.Down(20)
		case keyboard.Q:
			fmt.Println(key.Char)
			c.drone.Land()
		case keyboard.P:
			fmt.Println(key.Char)
			c.drone.TakeOff()
		case keyboard.ArrowUp:
			fmt.Println(key.Char)
			c.drone.Forward(20)
		case keyboard.ArrowDown:
			fmt.Println(key.Char)
			c.drone.Backward(20)
		case keyboard.ArrowLeft:
			fmt.Println(key.Char)
			c.drone.Left(20)
		case keyboard.ArrowRight:
			fmt.Println(key.Char)
			c.drone.Right(20)
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

}

// Down moves down the drone
func (c *KeyBoard) Down() {

}

// Right moves right the drone
func (c *KeyBoard) Right() {

}

// Left moves to left the drone
func (c *KeyBoard) Left() {

}

// TakeOff start the drone and go up it
func (c *KeyBoard) TakeOff() {
	fmt.Println("<-- TakeOff")
	c.drone.TakeOff()
}

// Land the drone
func (c *KeyBoard) Land() {
	fmt.Println("<-- Land")
	gobot.After(5*time.Second, func() {
		c.drone.Land()
	})
}
