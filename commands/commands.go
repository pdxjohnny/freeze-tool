package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/freeze-tool/connected"
	"github.com/pdxjohnny/freeze-tool/history"
	"github.com/pdxjohnny/freeze-tool/web"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "web",
		Short: "Web server and web socket server",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			web.Run()
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
		Short: "Keeps track of DeviceStatusUpdates",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			history.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
