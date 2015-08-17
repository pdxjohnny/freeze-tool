package adb

import (
	"log"
	"testing"
)

func TestDevices(t *testing.T) {
	log.Println(Devices())
}
