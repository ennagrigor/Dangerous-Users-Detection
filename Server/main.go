package main

import (
	"fmt"
	"userdetection/Configuration"
	"userdetection/Router"
	"userdetection/TextDetection"
)

const (
	configurationPath = `C:\Users\home\Desktop\UserDetection\Server\conf.json`
)

func main() {
	Configuration.InitConfiguration(configurationPath)
	TextDetection.ClassifyText(`I'm hero`)
	fmt.Println("Starting server")
	Router.InitServer()
}
