package main

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/freeze-tool/history"
	"github.com/pdxjohnny/freeze-tool/connected"
)

func main() {
	commands := []*cobra.Command{
		&cobra.Command{
			Use:   "web",
			Short: "Web server and web socket server",
			Run: func(cmd *cobra.Command, args []string) {
				ConfigBindFlags(cmd)
				Web()
			},
		},
		&cobra.Command{
			Use:   "connected",
			Short: "Watches for changes in plugged devices",
			Run: func(cmd *cobra.Command, args []string) {
				ConfigBindFlags(cmd)
				connected.Run()
			},
		},
		&cobra.Command{
			Use:   "history",
			Short: "Keeps track of DeviceStatus updates",
			Run: func(cmd *cobra.Command, args []string) {
				ConfigBindFlags(cmd)
				history.Run()
			},
		},
	}

	ConfigDefaults(commands...)
	var rootCmd = &cobra.Command{Use: "freeze-tool"}
	rootCmd.AddCommand(commands...)
	rootCmd.Execute()
}
