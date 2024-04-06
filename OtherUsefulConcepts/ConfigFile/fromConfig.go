package main

import (
	"fmt"

	"github.com/koding/multiconfig"
)

type Server struct {
	Name     string `required:"true"`
	Port     int    `default:"6060"`
	Enabled  bool
	Users    []string
	Database struct {
		Port int
	}
}

func main() {
	m := multiconfig.New()
	m = multiconfig.NewWithPath("config.toml")
	// Get an empty struct for your configuration
	serverConf := new(Server)

	// Populated the serverConf struct
	//err := m.Load(serverConf) // Check for error
	m.MustLoad(serverConf) // Panic's if there is any error

	// Access now populated fields
	fmt.Println(serverConf.Port) // by default 6060
	fmt.Println(serverConf.Name) // "koding"
}
