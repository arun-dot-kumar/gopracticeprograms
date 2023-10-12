package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	// Create a new scheduler
	s := gocron.NewScheduler(time.UTC)

	scheduler := s.CronWithSeconds("*/5 * * * * *")
	scheduler.ChangeLocation(loc)
	j, _ := scheduler.Do(func() {
		fmt.Println(time.Now().In(loc))
	})
	s.StartAsync()
	time.Sleep(20 * time.Second)
	fmt.Println("updating scheduled time")
	scheduler = s.Job(j).CronWithSeconds("*/6 * * * * *")
	scheduler.Update()
	time.Sleep(21 * time.Second)
}
