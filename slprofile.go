// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/carloct/slprofile/route"
	"github.com/carloct/slprofile/shared/jsonconfig"
	"github.com/carloct/slprofile/shared/server"
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
