package main

import (
	"log"

	"github.com/robfig/cron/v3"
	"github.com/wzhanjun/go-gin-example/models"
)

func main() {
	log.Println("Starting.....")

	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/1 * * * * *", func() {
		log.Println("Run models.CleanAllTag.....")
		models.CleanAllTag()
	})

	c.Start()

	// t1 := time.NewTicker(time.Second * 10)
	// for {
	// 	select {
	// 	case <-t1.C:
	// 		t1.Reset(time.Second * 10)
	// 	}
	// }
	select {}

}
