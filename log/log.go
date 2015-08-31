package log

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pdxjohnny/microsocket/service"
)

type Logger struct {
	*service.Service
	// A devices history
	History map[string][]byte
}

func NewLogger() *Logger {
	// Service setup
	inner := service.NewService()
	log := Logger{Service: inner}
	log.Caller = &log
	// Init Data map
	log.History = make(map[string][]interface{})
	return &log
}

func (log *Logger) DeviceStatus(raw_message []byte) {
	// Create a new message struct
	message := new(interface{})
	// Parse the message to a json
	err := json.Unmarshal(raw_message, &message)
	// Return if error or no DumpKey or not the client specified to dump
	if err != nil {
		return
	}
	// Otherwise
	// Check if this request is applicable to this instance
	_, ok := log.History[message["Device"]]
	// If it is then there will be a channel and this will
	if ok {
		// Send the response to the channel
		log.History[message["Device"]] <- true
	}
}
