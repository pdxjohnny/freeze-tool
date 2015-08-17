package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/pdxjohnny/easysock"

	"github.com/pdxjohnny/freeze-tool/config"
	"github.com/pdxjohnny/freeze-tool/connected"
)

type mainFunction func() error

func Web() error {
	conf := config.Load()
	go easysock.Hub.Run()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", easysock.ServeWs)
	port := fmt.Sprintf(":%s", conf.Port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Connected() error {
	conf := config.Load()
	wsUrl := fmt.Sprintf("http://%s:%s/ws", conf.Host, conf.Port)
	connectedService := connected.NewConnected()
	err := connectedService.Connect(wsUrl)
	if err != nil {
		log.Println(err)
	}
	connectedService.Read()
	return nil
}

func functionToRun() string {
	flag.Parse()
	runFunc := "web"
	args := flag.Args()
	if len(args) > 0 {
		runFunc = args[len(args)-1]
	}
	return strings.ToLower(runFunc)
}

func main() {
	runFunc := functionToRun()
	allFunctions := map[string]mainFunction{
		"web":       Web,
		"connected": Connected,
	}
	call, ok := allFunctions[runFunc]
	if ok {
		call()
	}
}
