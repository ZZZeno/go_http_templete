package middlewares

import (
	"github.com/robfig/cron/v3"
	"log"
)

var scheduler *cron.Cron

func registerJob1() cron.EntryID {
	entryId, _ := scheduler.AddFunc("@every 20s", func() {
		// do something ever 20s
	})
	return entryId
}

func SchedulerRegister() {
	go func() {
		scheduler = cron.New()
		registerJob1()
		scheduler.Start()
		log.Println("Schedulers are started.")
	}()
}
