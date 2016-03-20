package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	Hostname string `json:"Hostname"`
	HTTPPort int    `json:"HTTPPort"`
}

func Run(httpHandlers http.Handler, s Server) {
	startHTTP(httpHandlers, s)
}

func startHTTP(handlers http.Handler, s Server) {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.HTTPPort), handlers))
}
