package adb

import (
	"strings"
)

func Devices() ([]string, error) {
	cmdArgs := []string{"devices"}
	cmdOut, err := Command(cmdArgs...)
	if err != nil {
		return nil, err
	}
	adbOutput := string(cmdOut)
	rawDeviceList := strings.Split(adbOutput, "\n")
	numDevices := len(rawDeviceList) - 1
	deviceList := make([]string, 0)
	for listLength := 0; listLength < numDevices; listLength++ {
		item := rawDeviceList[listLength]
		// Only add usable devices to the list and don't add the first
		// line that says "List of devices attached"
		if strings.Contains(item, "device") &&
			!strings.Contains(item, "List") {
			deviceList = append(deviceList, strings.Split(item, "\t")[0])
		}
	}
	return deviceList, nil
}
