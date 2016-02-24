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
	case status.Disconnected, status.NotLogging, status.ErrorLogging:
		// If the device is no longer connected then stop logging
		// Also if there was an error logging or if we dont want to log
		if command.LogProcess == nil {
			return
		}
		err := command.LogProcess.Process.Kill()
		if err != nil {
			log.Println("Error killing logging process:", err)
		}
		command.LogProcess = nil

		// Send the current status of the logger
		command.Status = status.NotLogging
		command.SendInterface(command)
	default:
		// So long as the device was not just disconnected we should be logging
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
	}
}
