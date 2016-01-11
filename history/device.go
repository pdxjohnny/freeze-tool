package history

import (
	"encoding/json"
	"time"
)

// Dump history of device
type DumpMessage struct {
	Device   string
	Method   string
	DumpKey  string
	ClientId string
}

func (history *Historian) DumpTracker(DumpKey string) bool {
	// Allocate a channel
	_, ok := history.DumpTrack[DumpKey]
	if !ok {
		history.DumpTrack[DumpKey] = make(chan bool, 1)
	}
	// Wait for client of choose this service to Dump
	history.SendInterface(&DumpMessage{
		Method:   "SendDeviceHistoryConfirm",
		DumpKey:  DumpKey,
		ClientId: history.ClientId,
	})
	// If the service is not chosen in 5 seconds then it exits
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(5 * time.Second)
		timeout <- true
	}()
	select {
	case wasChosen := <-history.DumpTrack[DumpKey]:
		return wasChosen
	case <-timeout:
		return false
	}
	return false
}

func (history *Historian) SendDeviceHistory(raw_message []byte) {
	// Create a new message struct
	message := DumpMessage{}
	// Parse the message to a json
	err := json.Unmarshal(raw_message, &message)
	// Return if error or no DumpKey or no Device
	if err != nil || message.DumpKey == "" || message.Device == "" {
		return
	}
	// Make sure we have the devices past
	_, ok := history.Past[message.Device]
	if !ok {
		return
	}
	// Make sure the map is initialized
	shouldDump := history.DumpTracker(message.DumpKey)
	if !shouldDump {
		return
	}
	// Loop through all stored data
	for _, item := range history.Past[message.Device] {
		changeMethod, err := history.MapBytes([]byte(item))
		if err != nil {
			return
		}
		// Change the method to DeviceHistory
		changeMethod["Method"] = "DeviceHistory"
		// Add the DumpKey to the object
		changeMethod["DumpKey"] = message.DumpKey
		// Send the status update
		history.SendInterface(changeMethod)
	}
	// Done dumping no need to track dumped anymore
	delete(history.DumpTrack, message.DumpKey)
}

func (history *Historian) SendDeviceHistoryConfirmed(raw_message []byte) {
	// Create a new message struct
	message := DumpMessage{}
	// Parse the message to a json
	err := json.Unmarshal(raw_message, &message)
	// Return if error or no DumpKey or not the client specified to dump
	if err != nil || message.DumpKey == "" ||
		message.ClientId != history.ClientId {
		return
	}
	// Otherwise
	// Check if this request is applicable to this instance
	_, ok := history.DumpTrack[message.DumpKey]
	// If it is then there will be a channel and this will
	if ok {
		// Send the response to the channel
		history.DumpTrack[message.DumpKey] <- true
	}
}
