package leap

const (
	// StopStatus The circle gesture is finished.
	StopStatus = "stop"

	// StartStatus the circle gesture has just started. The movement has progressed far enough
	// for the recognizer to classify it as a circle.
	StartStatus = "start"

	// UpdateStatus The circle gesture is continuing.
	UpdateStatus = "update"
)
