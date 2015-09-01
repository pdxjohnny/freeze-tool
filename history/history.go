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
	log.Println("Listening for DeviceStatusUpdates...")
	history.Read()
	return nil
}

func (history *Historian) DeviceStatusUpdate(raw_message []byte) {
	// Create a new message struct
	message, err := history.MapBytes(raw_message)
	if err != nil {
		return
	}
	deviceName := message["Device"].(string)
	// Check if a slice needs to be made
	_, ok := history.Past[deviceName]
	if !ok {
		// Make a slice
		history.Past[deviceName] = make([]string, 0)
	}
	// Determine if this is different than the last status sent
	newStatus := false
	// Have to compare strings
	string_message := string(raw_message)
	// Index of last status
	lastIndex := len(history.Past[deviceName]) - 1
	// If the index is less than zero it is the first status
	if lastIndex < 0 {
		newStatus = true
	} else {
		// Otherwise compare to last status to see if its new
		last := history.Past[deviceName][lastIndex]
		if last != string_message {
			newStatus = true
		}
	}
	// Append to the devices history if
	// its not the same as the last info sent
	if newStatus {
		log.Println(deviceName, "was updated")
		history.Past[deviceName] = append(history.Past[deviceName], string_message)
		// Send this because it is the new status of the device
		history.SendDeviceStatus(deviceName)
	}
}

func (history *Historian) LastDeviceStatuses(raw_message []byte) {
	// Send the last status of every device
	for deviceName, _ := range history.Past {
		history.SendDeviceStatus(deviceName)
	}
}

func (history *Historian) SendDeviceStatus(deviceName string) {
	// Check if the device has a history
	_, ok := history.Past[deviceName]
	if !ok {
		return
	}
	// Send the last status for that device
	lastIndex := len(history.Past[deviceName]) - 1
	// If its greater than -1 ther is something in the array to send
	if lastIndex > -1 {
		last := history.Past[deviceName][lastIndex]
		// Change the method to DeviceStatus
		changeMethod, err := history.MapBytes([]byte(last))
		if err != nil {
			return
		}
		changeMethod["Method"] = "DeviceStatus"
		// Send the current status of the device
		history.SendInterface(changeMethod)
	}
}
