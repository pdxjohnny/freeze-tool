package adb

import (
	"testing"
)

func TestCommand(t *testing.T) {
  args := []string{"version"}
  _, err := Command(args...)
  if err != nil {
    panic("Error getting adb version")
  }
}
