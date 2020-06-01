package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"gobot.io/x/gobot/platforms/leap"
)

const (
	loggerFile = "/tmp/logger.log"
)

// Logger is an utility to show information related to different leap motion
// structures
type Logger interface {
	ShowGesture(g leap.Gesture)
	ShowHand(h leap.Hand, message ...string)
	ShowFingers(f leap.Frame)
}

type logger struct {
}

// NewLogger generates a pointer to logger.
func NewLogger() Logger {

	// Setup logger
	lumberjackLogRotate := &lumberjack.Logger{
		Filename:   loggerFile,
		MaxSize:    5,  // Max megabytes before log is rotated
		MaxBackups: 90, // Max number of old log files to keep
		MaxAge:     60, // Max number of days to retain log files
		Compress:   true,
	}

	mw := io.MultiWriter(os.Stdout, lumberjackLogRotate)
	log.SetOutput(mw)
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.JSONFormatter{})

	return &logger{}
}

// ShowGesture show the information related to gesture
func (l *logger) ShowGesture(g leap.Gesture) {
	log.WithFields(log.Fields{
		"type":     g.Type,
		"state":    g.State,
		"progress": g.Progress,
		"normal":   g.Normal,
	}).Info("gesture")
}

func (l *logger) ShowHand(h leap.Hand, message ...string) {

	info := fmt.Sprintf("hand:%s", message)
	log.WithFields(log.Fields{
		"id":   h.ID,
		"type": h.Type,
		"x":    h.X(),
		"y":    h.Y(),
		"z":    h.Z(),
	}).Info(info)
}

func (l *logger) ShowFingers(f leap.Frame) {
	ps := f.Pointables

	for _, p := range ps {
		log.WithFields(log.Fields{
			"id":      p.ID,
			"hand_id": p.HandID,
			"type":    p.Type,
		}).Info("finger")
	}
}
