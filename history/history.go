package history

import (
	"encoding/json"
	"log"

	"github.com/pdxjohnny/microsocket/service"
)

type DeviceStatusUpdate struct {
	Device string
}

type Historian struct {
	*service.Service
	// History of status updates for each device
	Past map[string][]string
}

func NewHistorian() *Historian {
	// Service setup
	inner := service.NewService()
	history := Historian{Service: inner}
	history.Caller = &history
	// Init Data map
	history.Past = make(map[string][]string)
	return &history
}

func (history *Historian) Run() error {
	log.Println("Listening to DeviceStatus updates...")
	history.Read()
	return nil
}

func (history *Historian) DeviceStatus(raw_message []byte) {
	// Create a new message struct
	message := DeviceStatusUpdate{}
	// Parse the message to a json
	err := json.Unmarshal(raw_message, &message)
	// Return if error or no DumpKey or not the client specified to dump
	if err != nil {
		return
	}
	// Otherwise
	log.Println(message.Device, "was updated")
	// Check if a slice needs to be made
	_, ok := history.Past[message.Device]
	if !ok {
		// Make a slice
		history.Past[message.Device] = make([]string, 0)
	}
	// Append to the devices history if
	// its not the same as the last info sent
	newStatus := false
	string_message := string(raw_message)
	lastIndex := len(history.Past[message.Device]) - 1
	if lastIndex < 0 {
		newStatus = true
	} else {
		last := history.Past[message.Device][lastIndex]
		if last != string_message {
			newStatus = true
		}
	}
	if newStatus {
		history.Past[message.Device] = append(history.Past[message.Device], string_message)
		log.Println(history.Past[message.Device])
	}
}
