package connected

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Run() error {
	wsURL := fmt.Sprintf("ws://%s:%d/ws",
		viper.GetString("host"),
		viper.GetInt("port"),
	)
	connected := NewConnected()
	log.Println("Connecting to", wsURL)
	err := connected.Connect(wsURL)
	if err != nil {
		log.Println(err)
	}
	connected.Run()
	return nil
}
