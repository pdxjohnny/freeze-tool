package connected

import (
	"log"
	"os"
	"time"

	"github.com/pdxjohnny/microsocket/service"
	stringArray "github.com/pdxjohnny/strings"
	"github.com/rjeczalik/notify"

	"github.com/pdxjohnny/freeze-tool/adb"
)

// Connected is a service which keeps teh connected devices updated
type Connected struct {
	*service.Service
	Hostname      string
	OldDeviceList map[string]bool
}

// NewConnected creates a new connected service
func NewConnected() *Connected {
	// Service setup
	inner := service.NewService()
	connected := Connected{Service: inner}
	connected.Caller = &connected
	connected.Hostname, _ = os.Hostname()
	connected.OldDeviceList = make(map[string]bool)
	return &connected
}

// Run starts the connected service
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

// SendHostname sends this computers hostname so its id can be
// associated with it
func (connected *Connected) SendHostname(rawMessage []byte) {
	// The message containing devices connect to the host
	connected.SendInterface(map[string]interface{}{
		"Method": "Hostname",
		"Name":   connected.ClientId,
		"Host":   connected.Hostname,
	})
}

// SendDevices sends the list of connected devices via adb
func (connected *Connected) SendDevices(rawMessage []byte) {
	log.Println("SendDevices")
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

// WatchDeviceChange starts a file system watcher to watch for
// usb device changes
func (connected *Connected) WatchDeviceChange() {
	// Make the channel buffered to ensure no event is dropped. Notify will drop
	// an event if the receiver is not able to keep up the sending pace.
	c := make(chan notify.EventInfo, 1)
	defer notify.Stop(c)

	// Set up a watchpoint listening for events within a directory tree rooted
	// at current working directory. Dispatch remove events to c.
	err := notify.Watch("/dev/bus/usb/...", c, notify.Create, notify.Remove)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-c:
			log.Println(event)
			go func() {
				time.Sleep(1500 * time.Millisecond)
				connected.SendDevices(nil)
			}()
		}
	}
}

// CheckOldDevices lets us knwo if any devices have been disconnected
func (connected *Connected) CheckOldDevices(deviceList []string) {
	for item := range connected.OldDeviceList {
		if !stringArray.Contains(deviceList, item) {
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
