package controller

// EventHandler contain the information releated to one event
// related to drone controller
type EventHandler struct {

	// EventName contain the name in which should trigger the handler function
	EventName string

	// Function is the function to execute on specific event
	Function func(data interface{})
}
