package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pdxjohnny/easysock"
	"github.com/spf13/viper"
)

func Web() error {
	go easysock.Hub.Run()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", easysock.ServeWs)
	port := fmt.Sprintf(":%d", viper.GetInt("port"))
	log.Println("Serving on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
