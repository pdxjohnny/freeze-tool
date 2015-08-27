package connected

import (
	"encoding/json"
	"os"
	"strings"
	"path/filepath"

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
	go connected.WatchDeviceChange()
	connected.SendDevices(nil)
	connected.Read()
	return nil
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
		"Host": connected.Hostname,
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
	recurse := true // include all sub directories
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
