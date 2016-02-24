package logger

import (
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

	log.Println("deviceStatus:", deviceStatus)

	// Take action based on the status of the device we are monitoring
	switch deviceStatus {
	case status.Disconnected, status.NotLogging, status.LoggingError:
		// If the device is no longer connected then stop logging
		// Also if there was an error logging or if we dont want to log
		if command.LogProcess == nil {
			return
		}
		err := command.LogProcess.Kill()
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

		cmd := exec.Command(
			"adb",
			"-s",
			command.Device,
			command.Command,
		)
		command.LogProcess = &Output{
			Cmd:  cmd,
			File: command.File,
		}
		err := command.LogProcess.Start()
		if err != nil {
			command.Status = status.LoggingError
			command.Message = err.Error()
			command.SendInterface(command)
			return
		}

		// Send the current status of the logger
		command.Status = status.Logging
		command.SendInterface(command)
		return
	}
}
