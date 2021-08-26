package main

import (
	"context"
	"os"
	"path"
	"time"
	"userdetection/Configuration"
	"userdetection/Database"
	"userdetection/Router"
	"userdetection/Scheduler"
	"userdetection/TextDetection"
)

const (
	configurationFile = "conf.json"
)

func initServices() {
	err := Configuration.InitConfiguration(configurationFile)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(Configuration.Conf.ApplicationFolder)
	if os.IsNotExist(err) {
		dirErr := os.Mkdir(Configuration.Conf.ApplicationFolder, os.ModePerm)
		if dirErr != nil {
			panic(dirErr)
		}
	}

	Configuration.Conf.IndexPath = path.Join(Configuration.Conf.ApplicationFolder, Database.IndexName)
	err = Database.InitIndex()
	if err != nil {
		panic(err)
	}
}

func initScheduledJobs() {
	scheduler := Scheduler.NewScheduler()
	scheduler.Add(context.Background(), TextDetection.DetectDangerousUserTweet, time.Minute*5)
}

func main() {
	initServices()
	initScheduledJobs()
	Router.InitServer()
}
