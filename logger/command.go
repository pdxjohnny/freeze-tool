package logger

import (
	"log"

	"github.com/pdxjohnny/microsocket/service"
)

type Command struct {
	*service.Service
	// Device is the device id we are connected to
	Device string `json:"device"`
	// Status lets clients know if the logging stream is active
	Status string `json:"status"`
	// Command we are running to collect logs (dmesg, logcat)
	Command string `json:"command"`
	// Method for when we send a status update. It will jsut send itself
	Method string `json:"method"`
	// File we are logging to and streaming from
	File string `json:"-"`
}

func NewCommand() *Command {
	// Service setup
	inner := service.NewService()
	command := Command{
		Service: inner,
		Method:  "DeviceStatusUpdate",
	}
	command.Caller = &command
	return &command
}

func (command *Command) Run() error {
	log.Println("Listening for DeviceStatusUpdates...")
	command.Read()
	return nil
}
