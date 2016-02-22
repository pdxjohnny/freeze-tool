package logger

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func Run() {
	logger := NewLogger()
	logger.Device = viper.GetString("device")
	logger.LogDir = viper.GetString("logdir")
	logger.Host = viper.GetString("host")
	logger.Port = viper.GetInt("port")

	mux := logger.Create()

	port := fmt.Sprintf(":%d", viper.GetInt("hport"))
	log.Println("Serving on", port)

	log.Fatal(http.ListenAndServe(port, mux))
}
