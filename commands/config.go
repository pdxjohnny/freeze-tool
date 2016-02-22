package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigOptions = map[string]interface{}{
	"hport": map[string]interface{}{
		"value": 7777,
		"help":  "Port to host on",
	},
	"port": map[string]interface{}{
		"value": 7777,
		"help":  "Port to connect to",
	},
	"host": map[string]interface{}{
		"value": "localhost",
		"help":  "Service host to connect websocket to",
	},
	"address": map[string]interface{}{
		"value": "0.0.0.0",
		"help":  "Address to host on",
	},
	"device": map[string]interface{}{
		"value": "",
		"help":  "Device id",
	},
	"logdir": map[string]interface{}{
		"value": "tmp",
		"help":  "Directory to store logs in",
	},
}

func ConfigDefaults(cmdList ...*cobra.Command) {
	ConfigEnv()
	ConfigSet()
	ConfigFlags(cmdList...)
}

func ConfigSet() {
	for index, item := range ConfigOptions {
		opt := item.(map[string]interface{})
		viper.SetDefault(index, opt["value"])
	}
}

func ConfigFlags(cmdList ...*cobra.Command) {
	for _, cmd := range cmdList {
		for index, item := range ConfigOptions {
			opt := item.(map[string]interface{})
			help := opt["help"].(string)
			switch value := opt["value"].(type) {
			case int:
				cmd.Flags().Int(index, value, help)
			case bool:
				cmd.Flags().Bool(index, value, help)
			case string:
				cmd.Flags().String(index, value, help)
			default:
			}
		}
	}
}

func ConfigBindFlags(cmd *cobra.Command) {
	for index, _ := range ConfigOptions {
		viper.BindPFlag(index, cmd.Flags().Lookup(index))
	}
}

func ConfigEnv() {
	viper.SetEnvPrefix("freezetool")
	viper.AutomaticEnv()
}
