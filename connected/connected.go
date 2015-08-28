package connected

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/andreaskoch/go-fswatch"
	"github.com/pdxjohnny/microsocket/service"

	"github.com/pdxjohnny/freeze-tool/adb"
)

type Connected struct {
	*service.Service
	Hostname string
}

func NewConnected() *Connected {
	// Service setup
	inner := service.NewService()
	connected := Connected{Service: inner}
	connected.Caller = &connected
	connected.Hostname, _ = os.Hostname()
	return &connected
}

func (connected *Connected) Run() error {
	go connected.Read()
	log.Println("Waiting for ready")
	<-connected.Ready
	log.Println("Sending devices...")
	connected.SendDevices(nil)
	log.Println("Sending hostname...")
	connected.SendHostname(nil)
	log.Println("Watching for change")
	connected.WatchDeviceChange()
	return nil
}

func (connected *Connected) SendHostname(raw_message []byte) {
	// The message containing devices connect to the host
	message := map[string]interface{}{
		"Method": "Hostname",
		"Name":   connected.ClientId,
		"Host":   connected.Hostname,
	}
	// Turn the message into json bytes
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return
	}
	// Dump it to clients
	connected.Write(messageBytes)
}

func (connected *Connected) SendDevices(raw_message []byte) {
	// Get the connected devices
	deviceList, err := adb.Devices()
	if err != nil {
		return
	}
	log.Println("Devices changed", deviceList)
	// Make the status for each device
	deviceStatus := make(map[string]interface{})
	for _, item := range deviceList {
		currentDevice := map[string]interface{}{
			"Device": item,
			"Host":   connected.ClientId,
			"Status": "Connected",
		}
		deviceStatus[item] = &currentDevice
	}
	// The message containing devices connect to the host
	message := map[string]interface{}{
		"Method":  "Devices",
		"Devices": deviceStatus,
		"Name":    connected.ClientId,
	}
	// Turn the message into json bytes
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return
	}
	// Dump it to clients
	connected.Write(messageBytes)
}

func (connected *Connected) WatchDeviceChange() {
	recurse := true
	skipDotFilesAndFolders := func(path string) bool {
		return strings.HasPrefix(filepath.Base(path), ".")
	}
	checkIntervalInSeconds := 1

	folderWatcher := fswatch.NewFolderWatcher(
		"/dev/bus/usb",
		recurse,
		skipDotFilesAndFolders,
		checkIntervalInSeconds,
	)
	folderWatcher.Start()

	for folderWatcher.IsRunning() {
		select {
		case <-folderWatcher.Modified():
			connected.SendDevices(nil)
		}
	}
}
