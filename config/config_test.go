package config

import (
	"log"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	conf := Load()
	log.Println(conf)
}
