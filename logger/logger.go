package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"path"

	"github.com/pdxjohnny/freeze-tool/status"
)

const (
	restError   = "{\"Error\": \"%s\"}"
	restSuccess = "{\"OK\": \"%s\"}"
)

type Logger struct {
	Device   string              `json:"device"`
	LogDir   string              `json:"logdir"`
	Host     string              `json:"host"`
	Port     int                 `json:"port"`
	Commands map[string]*Command `json:"-"`
}

func NewLogger() *Logger {
	logger := &Logger{
		Commands: make(map[string]*Command, 0),
	}
	return logger
}

func (logger *Logger) Create() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/status/", logger.needCommand(http.HandlerFunc(logger.statusHandler)))
	mux.Handle("/create/", logger.needCommand(http.HandlerFunc(logger.createHandler)))
	mux.Handle("/log/", logger.needCommand(http.HandlerFunc(logger.logHandler)))

	return mux
}

func (logger *Logger) needCommand(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		queryParams := u.Query()

		_, ok := queryParams["command"]
		if !ok {
			logger.Error(w, errors.New("Need \"command\" query var"), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (logger *Logger) statusHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	queryParams := u.Query()

	// If the command is running return its status
	command, ok := logger.Commands[queryParams["command"][0]]
	if ok {
		// Send the logger command that is already running back to the client
		dump := json.NewEncoder(w)
		logger.Error(w, dump.Encode(command), http.StatusInternalServerError)
		return
	}

	// The command is not running see if the client wanted the status of the
	// server
	if queryParams["command"][0] == "_self" {
		// Send the logger server info
		dump := json.NewEncoder(w)
		logger.Error(w, dump.Encode(logger), http.StatusInternalServerError)
		return
	}

	// We dont have what the client wants, sorry
	logger.Error(w, errors.New("Command not found"), http.StatusNotFound)
}

func (logger *Logger) createHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	queryParams := u.Query()

	// Check to make sure this command is not already running
	command, ok := logger.Commands[queryParams["command"][0]]
	if ok {
		log.Println("Already exists", queryParams["command"][0])
		// Send the logger command that is already running back to the client
		dump := json.NewEncoder(w)
		logger.Error(w, dump.Encode(command), http.StatusInternalServerError)
		return
	}

	log.Println("Creating new", queryParams["command"][0])
	// Create the command
	command = NewCommand()
	// The device is always the same for each logger service
	// this is because it is easier to debug memory and cpu usage
	// if there is one server per device
	command.Device = logger.Device
	command.Status = status.LoggingCreated
	command.Command = queryParams["command"][0]
	command.File = path.Join(
		logger.LogDir,
		command.Device,
		command.Command,
		".log",
	)

	// Format the ws server URL
	wsUrl := fmt.Sprintf("ws://%s:%d/ws",
		logger.Host,
		logger.Port,
	)
	// Connect the logger to the ws server
	err = command.Connect(wsUrl)
	if err != nil {
		logger.Error(w, err, http.StatusServiceUnavailable)
		return
	}

	// Send the current status of the logger
	command.SendInterface(command)
	// Start the logger command it will now listen to events from the ws server
	go command.Run()

	// Add the command to the map of already running commands
	logger.Commands[command.Command] = command

	// Send the logger command we just started back to the client
	dump := json.NewEncoder(w)
	logger.Error(w, dump.Encode(command), http.StatusInternalServerError)
}

func (logger *Logger) logHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	queryParams := u.Query()

	// If the command is running return its status
	command, ok := logger.Commands[queryParams["command"][0]]
	if !ok {
		// We dont have what the client wants, sorry
		logger.Error(w, errors.New("Command not found"), http.StatusNotFound)
		return
	}

	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}

	cmd := exec.Command(
		"tail",
		"-f",
		command.File,
	)
	cmd.Stdout = &fw
	cmd.Stderr = &fw
	cmd.Run()
}

func (logger *Logger) Error(w http.ResponseWriter, err error, status int) {
	if err == nil {
		return
	}
	w.WriteHeader(status)
	fmt.Fprintln(w, fmt.Sprintf(restError, err.Error()))
}

func (logger *Logger) Success(w http.ResponseWriter, message string) {
	fmt.Fprintln(w, fmt.Sprintf(restSuccess, message))
}

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}
