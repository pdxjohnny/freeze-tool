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
	connected := NewConnected()
	log.Println("Connecting to", wsUrl)
	err := connected.Connect(wsUrl)
	if err != nil {
		log.Println(err)
	}
	connected.Run()
	return nil
}
