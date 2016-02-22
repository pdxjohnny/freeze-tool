package logger

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/pdxjohnny/freeze-tool/status"
)

// Send DeviceStatusUpdate when we start logging

func (command *Command) DeviceStatus(raw_message []byte) {
	// Create a new message struct
	message, err := command.MapBytes(raw_message)
	if err != nil {
		return
	}
	deviceName := message["Device"].(string)
	deviceStatus := message["Status"].(string)

	// Only check logging on updates to our device
	if deviceName != command.Device {
		return
	}

	log.Println("Got an update to device", command.Device, deviceStatus)

	// Take action based on the status of the device we are monitoring
	switch deviceStatus {
	case status.Connected, status.LoggingCreated:
		if command.LogProcess != nil {
			return
		}
		pipeCommand := fmt.Sprintf(
			"adb -s %s %s &>%s",
			command.Device,
			command.Command,
			command.File,
		)
		log.Println(pipeCommand)
		command.LogProcess = exec.Command(
			"bash",
			"-c",
			pipeCommand,
		)
		command.LogProcess.Start()

		// Send the current status of the logger
		command.Status = status.Logging
		command.SendInterface(command)
	case status.Disconnected:
		if command.LogProcess == nil {
			return
		}
		err := command.LogProcess.Process.Kill()
		log.Println(err)
		command.LogProcess = nil
	}
}
