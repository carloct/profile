package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"

	"github.com/carloct/slprofile/route"
	"github.com/carloct/slprofile/shared/jsonconfig"
	"github.com/carloct/slprofile/shared/server"
	"github.com/carloct/slprofile/shared/view"

	"github.com/joho/godotenv"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	//Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root)

	server.Run(route.LoadHTTP(), config.Server)
}

// Application configuration
var config = &configuration{}

type configuration struct {
	Server   server.Server `json:"Server"`
	View     view.View     `json:"View"`
	Template view.Template `json:"Template"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
