package main

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/freeze-tool/connected"
)

func main() {
	var cmdWeb = &cobra.Command{
		Use:   "web",
		Short: "Web server and web socket server",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			Web()
		},
	}

	var cmdConnected = &cobra.Command{
		Use:   "connected",
		Short: "Watches for changes in plugged devices",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			connected.Run()
		},
	}

	ConfigDefaults(
		cmdWeb,
		cmdConnected,
	)
	var rootCmd = &cobra.Command{Use: "freeze-tool"}
	rootCmd.AddCommand(
		cmdWeb,
		cmdConnected,
	)
	rootCmd.Execute()
}
