package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"

	"github.com/carloct/jsproducts/route"
	"github.com/carloct/jsproducts/shared/jsonconfig"
	"github.com/carloct/jsproducts/shared/server"

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

	server.Run(route.LoadHTTP(), config.Server)
}

// Application configuration
var config = &configuration{}

type configuration struct {
	Server server.Server `json:"Server"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
