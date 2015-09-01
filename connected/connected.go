package connected

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/andreaskoch/go-fswatch"
	string_array "github.com/pdxjohnny/strings"
	"github.com/pdxjohnny/microsocket/service"

	"github.com/pdxjohnny/freeze-tool/adb"
)

type Connected struct {
	*service.Service
	Hostname string
	OldDeviceList map[string]bool
}

func NewConnected() *Connected {
	// Service setup
	inner := service.NewService()
	connected := Connected{Service: inner}
	connected.Caller = &connected
	connected.Hostname, _ = os.Hostname()
	connected.OldDeviceList = make(map[string]bool)
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
	connected.SendInterface(map[string]interface{}{
		"Method": "Hostname",
		"Name":   connected.ClientId,
		"Host":   connected.Hostname,
	})
}

func (connected *Connected) SendDevices(raw_message []byte) {
	// Get the connected devices
	deviceList, err := adb.Devices()
	if err != nil {
		return
	}
	connected.CheckOldDevices(deviceList)
	log.Println("Devices changed", deviceList)
	for _, item := range deviceList {
		currentDevice := map[string]interface{}{
			"Method": "DeviceStatusUpdate",
			"Device": item,
			"Host":   connected.ClientId,
			"Status": "Connected",
		}
		connected.SendInterface(currentDevice)
	}
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

func (connected *Connected) CheckOldDevices(deviceList []string) {
	for item, _ := range connected.OldDeviceList {
		if !string_array.Contains(deviceList, item) {
			currentDevice := map[string]interface{}{
				"Method": "DeviceStatusUpdate",
				"Device": item,
				"Host":   connected.ClientId,
				"Status": "Disconnected",
			}
			connected.SendInterface(currentDevice)
			delete(connected.OldDeviceList, item)
		}
	}
	for _, item := range deviceList {
		_, ok := connected.OldDeviceList[item]
		if !ok {
			connected.OldDeviceList[item] = true
		}
	}
}
