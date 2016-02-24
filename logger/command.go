package logger

import (
	"log"

	"github.com/pdxjohnny/microsocket/service"
)

type Command struct {
	*service.Service
	// Device is the device id we are connected to
	Device string `json:"Device"`
	// Status lets clients know if the logging stream is active
	Status string `json:"Status"`
	// Message lets us know extra information beyond status
	Message string `json:"Message"`
	// Command we are running to collect logs (dmesg, logcat)
	Command string `json:"Command"`
	// Method for when we send a status update. It will jsut send itself
	Method string `json:"Method"`
	// File we are logging to and streaming from
	File string `json:"-"`
	// LogProcess is the adb process capturing log output
	LogProcess *Output `json:"-"`
}

func NewCommand() *Command {
	// Service setup
	inner := service.NewService()
	command := Command{
		Service:    inner,
		Method:     "DeviceStatusUpdate",
		LogProcess: nil,
	}
	command.Caller = &command
	return &command
}

func (command *Command) Run() error {
	log.Println("Listening for DeviceStatusUpdates...")
	command.Read()
	return nil
}
