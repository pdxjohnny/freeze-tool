package connected

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Run() error {
	wsUrl := fmt.Sprintf("http://%s:%d/ws",
		viper.GetString("host"),
		viper.GetInt("port"),
	)
	log.Println("Connecting to", wsUrl)
	connected := NewConnected()
	err := connected.Connect(wsUrl)
	if err != nil {
		log.Println(err)
	}
	connected.Run()
	return nil
}
