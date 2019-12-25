package controller

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// LeapMotion is a controller in charge to handle the
// moviements of tello drone with LeapMotion device.
type LeapMotion struct {
	drone *tello.Driver
	leap  *leap.Driver
}

// NewLeapMotion generates pointer to LeapMotion
func NewLeapMotion(drone *tello.Driver, leap *leap.Driver) *LeapMotion {
	return &LeapMotion{drone: drone, leap: leap}
}

// TakeOff the drone with the
func (l *LeapMotion) TakeOff() {

}

func (l *LeapMotion) Up()    {}
func (l *LeapMotion) Down()  {}
func (l *LeapMotion) Right() {}
func (l *LeapMotion) Left()  {}

func (l *LeapMotion) Land() {}
func (l *LeapMotion) Run()  {}
