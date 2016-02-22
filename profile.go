package main

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/carloct/profile/route"
	"github.com/carloct/profile/shared/database"
	"github.com/carloct/profile/shared/jsonconfig"
	"github.com/carloct/profile/shared/server"
	"github.com/carloct/profile/shared/session"
	"github.com/carloct/profile/shared/view"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	session.Configure(config.Session)
	database.Connect(config.Database)

	view.Configure(config.View)
	server.Run(route.LoadHTTP(), config.Server)
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Server   server.Server      `json:"Server"`
	Template view.Template      `json:"Template"`
	View     view.View          `json:"View"`
	Session  session.Session    `json:"Session"`
	Database database.Databases `json:"Database"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
