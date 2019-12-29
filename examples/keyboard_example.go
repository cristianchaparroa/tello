package main

import (
	"tello/controller/keyboard"

	"gobot.io/x/gobot"
	tello "gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8888")
	c := keyboard.NewController(drone)

	work := func() {
		c.Run()
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{c.GetKeyBoard(), drone},
		work,
	)

	robot.Start()
}
