package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pdxjohnny/microsocket/server"
	"github.com/spf13/viper"
)

func Run() error {
	go server.Hub.Run()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", server.Hub.ServeWs)
	port := fmt.Sprintf(":%d", viper.GetInt("port"))
	log.Println("Serving on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
