package main

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/freeze-tool/commands"
)

func main() {
	var rootCmd = &cobra.Command{Use: "freeze-tool"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
