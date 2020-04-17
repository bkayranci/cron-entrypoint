package main

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/robfig/cron/v3"
	"os/exec"
	"sync"
)

var (
	configCron       = kingpin.Flag("cron.schedule", "Cron schedule").Default("* * * * *").String()
	executable       = kingpin.Flag("cron.executable", "Cron executable command").Required().String()
	execOnStart		 = kingpin.Flag("exec.onstart", "Execute one more time on started").Bool()
	handleException  = kingpin.Flag("handle.exception", "Handle via log error on exception").Bool()
)

func execute() {
	result, err := exec.Command(*executable).Output()
	if err != nil {
		if *handleException == true {
			log.Error(err)
		} else {
			log.Fatal(err)
		}
	}

	log.Infof("Stdout: %s", result)
}

func main() {
	kingpin.Parse()
	log.Infof("cron: %s", *configCron)
	log.Infof("executable: %s", *executable)

	/**
	 * Execute on more time on start
	 */
	if *execOnStart == true {
		execute()
	}

	var wg sync.WaitGroup
	wg.Add(1)

	c := cron.New()
	c.AddFunc(*configCron, execute)
	c.Start()

	wg.Wait()
}
