package history

import (
	"log"

	"github.com/pdxjohnny/microsocket/service"
)

const (
	DefaultMaxLength = 20
)

type DeviceMessage struct {
	Device string
}

type Historian struct {
	*service.Service
	// MaxLength is the length of history each device will store
	// This prevents using to much memory storing device connected
	// device disconnected a million times
	MaxLength int
	// History of status updates for each device
	Past map[string][]string
	// Keep track of if we should dump or not
	DumpTrack map[string]chan bool
}

func NewHistorian() *Historian {
	// Service setup
	inner := service.NewService()
	history := Historian{
		Service:   inner,
		MaxLength: DefaultMaxLength,
		// Init Past map
		Past: make(map[string][]string),
		// Init DumpTrack map
		DumpTrack: make(map[string]chan bool),
	}
	history.Caller = &history
	return &history
}

func (history *Historian) Run() error {
	log.Println("Listening for DeviceStatusUpdates...")
	history.Read()
	return nil
}
