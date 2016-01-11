package history

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Run() error {
	wsUrl := fmt.Sprintf("ws://%s:%d/ws",
		viper.GetString("host"),
		viper.GetInt("port"),
	)
	history := NewHistorian()
	log.Println("Connecting to", wsUrl)
	err := history.Connect(wsUrl)
	if err != nil {
		log.Println(err)
	}
	history.Run()
	return nil
}
