package connected

import (
	"encoding/json"

	"github.com/pdxjohnny/microsocket/service"

	"github.com/pdxjohnny/freeze-tool/adb"
)

type Connected struct {
	*service.Service
}

func NewConnected() *Connected {
	// Service setup
	inner := service.NewService()
	connected := Connected{Service: inner}
	connected.Caller = &connected
	return &connected
}

func (connected *Connected) SendDevices(raw_message []byte) {
	// Get the connected devices
	deviceList, err := adb.Devices()
	if err != nil {
		return
	}
	// The message containing devices connect to the host
	message := map[string]interface{}{
		"Method":  "Devices",
		"Devices": deviceList,
	}
	// Turn the message into json bytes
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return
	}
	// Dump it to clients
	connected.Write(messageBytes)
}
