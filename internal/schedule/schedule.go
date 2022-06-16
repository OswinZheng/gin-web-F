package schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

/**
定时任务
*/

type testJob struct {
}

func (t testJob) Run() {
	fmt.Println("test job")
}

func Start() {
	c := cron.New()

	c.AddFunc("@every 1s", func() {
		fmt.Println("test fun")
	})

	c.AddJob("@every 1s", testJob{})

	c.Start()

	select {}
}
