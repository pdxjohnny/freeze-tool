package logger

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/pdxjohnny/diffstream/diff"
)

type Output struct {
	Diff *diff.Diff
	Cmd  *exec.Cmd
	File string
}

func (output *Output) Start() error {
	outputFile, err := os.OpenFile(
		output.File,
		os.O_APPEND|os.O_WRONLY|os.O_CREATE,
		0600,
	)
	if err != nil {
		return err
	}

	intputFile, err := os.OpenFile(
		output.File,
		os.O_RDONLY,
		0600,
	)
	if err != nil {
		return err
	}

	cmdOutputReader, cmdOutputWriter := io.Pipe()

	output.Cmd.Stdout = cmdOutputWriter
	output.Cmd.Stderr = cmdOutputWriter

	output.Diff = &diff.Diff{
		First:  intputFile,
		Second: cmdOutputReader,
		Output: outputFile,
	}

	output.Cmd.Start()
	go output.Diff.Start()
	return nil
}

func (output *Output) Kill() error {
	err := output.Diff.First.(*os.File).Close()
	if err != nil {
		log.Println(err)
	}
	err = output.Diff.Output.(*os.File).Close()
	if err != nil {
		log.Println(err)
	}
	return output.Cmd.Process.Kill()
}
