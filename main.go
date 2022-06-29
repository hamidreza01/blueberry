package main

import (
	"runtime"

	"github.com/hamidreza01/blueberry/config"
	"github.com/hamidreza01/blueberry/network"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	network.Start(config.Config{
		Ip:       "127.0.0.1",
		Port:     ":1000",
		Error404: "",
		Error500: "",
		Error403: "",
		Error203: "",
		Error:    "./data/error.html",
	})
}
