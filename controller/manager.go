package controller

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

// Manager is in charge to handle multiple tello controllers
type Manager struct {
	drone       *tello.Driver
	Controllers []Controller
}

// NewManager returns a pointer to Manager
func NewManager(cs ...Controller) *Manager {
	return &Manager{Controllers: cs}
}

// Add register a controller to be handled
func (m *Manager) Add(c Controller) {
	m.Controllers = append(m.Controllers, c)
}

// GetListeners returns the listeners for diferentes devices
func (m *Manager) GetListeners() []*EventHandler {

	events := make([]*EventHandler, 0)
	for _, c := range m.Controllers {
		es := c.GetEvents()
		events = append(events, es...)
	}
	return events
}

// GetDevices retrieves all the devices to be managed
func (m *Manager) GetDevices() []gobot.Device {
	devices := make([]gobot.Device, 0)

	return devices
}

// Start inits the different devices
func (m *Manager) Start() {

	events := m.GetListeners()
	work := func() {
		for _, e := range events {
			m.drone.On(e.EventName, e.Function)
		}
	}

	devices := m.GetDevices()
	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		devices,
		work,
	)

	robot.Start()
}
