package main

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/carloct/profile/route"
	"github.com/carloct/profile/shared/jsonconfig"
	"github.com/carloct/profile/shared/server"
)

func init() {
	// Verbose logging with file name and line number
	//log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	server.Run(route.LoadHTTP(), config.Server)
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Server server.Server `json:"Server"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
