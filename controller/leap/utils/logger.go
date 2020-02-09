package utils

import (
	"fmt"
	"strings"

	"gobot.io/x/gobot/platforms/leap"
)

// Logger is an utility to show information related to different leap motion
// structures
type Logger interface {
	ShowGesture(g leap.Gesture)
	ShowHand(h leap.Hand)
	ShowFingers(f leap.Frame)
}

type logger struct {
}

// NewLogger generates a pointer to logger.
func NewLogger() Logger {
	return &logger{}
}

// ShowGesture show the information related to gesture
func (l *logger) ShowGesture(g leap.Gesture) {
	message := fmt.Sprintf( "Gesture -> Type:%s, State:%s, IsInProggress:%v, Normal:%v", g.Type, g.State, g.Progress, g.Normal)
	fmt.Println(message)
}

func (l *logger) ShowHand(h leap.Hand) {
	message := fmt.Sprintf("Hand -> \n")
	fmt.Println(message)
}

func (l *logger) ShowFingers(f leap.Frame) {
	var message strings.Builder
	ps := f.Pointables

	for _, p := range ps {
		ms := fmt.Sprintf("--> Finger ID:%v, HandID:%v, type:%v \n", p.ID, p.HandID, p.Type)
		message.WriteString(ms)
	}
	fmt.Println(message.String())
}