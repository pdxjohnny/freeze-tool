package adb

import (
	"os/exec"
)

const (
	// The command line program to execute
	cmdName = "adb"
)

func Command(arguments ...string) (string, error) {
	cmdOut, err := exec.Command(cmdName, arguments...).Output()
	if err != nil {
		return "", err
	}
	return string(cmdOut), nil
}
