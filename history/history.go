package history

import (
	"log"

	"github.com/pdxjohnny/microsocket/service"
)

type DeviceMessage struct {
	Device string
}

type Historian struct {
	*service.Service
	// History of status updates for each device
	Past map[string][]string
	// Keep track of if we should dump or not
	DumpTrack map[string]chan bool
}

func NewHistorian() *Historian {
	// Service setup
	inner := service.NewService()
	history := Historian{Service: inner}
	history.Caller = &history
	// Init Past map
	history.Past = make(map[string][]string)
	// Init DumpTrack map
	history.DumpTrack = make(map[string]chan bool)
	return &history
}

func (history *Historian) Run() error {
	log.Println("Listening for DeviceStatusUpdates...")
	history.Read()
	return nil
}
