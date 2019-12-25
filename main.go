package main

import (
	"fmt"
	"tello/controller"

	"gobot.io/x/gobot"
	tello "gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

func main() {
	drone := tello.NewDriver("8888")
	keys := keyboard.NewDriver()
	c := controller.NewKeyBoard(drone, keys)

	var flightData *tello.FlightData
	work := func() {
		drone.On(tello.FlightDataEvent, func(data interface{}) {
			flightData = data.(*tello.FlightData)
			fmt.Println("Height:", flightData.Height)
		})
		c.Run()
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{keys, drone},
		work,
	)

	robot.Start()
}
