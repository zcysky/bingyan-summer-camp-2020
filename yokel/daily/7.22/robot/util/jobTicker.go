package util

import (
	"fmt"
	"github.com/Logiase/gomirai"
	"time"
)

type JobTicker struct {
	FinishJob func(ticker *JobTicker)
	Job         func(ticker *JobTicker)
	NotiId      int64
	Interval    int64                   `json:"interval"`
	TargetTime  int64                   `json:"target-time"`
	TargetId    uint                    `json:"target-id"`
	TargetEvent int64                   `json:"target-event"`
	Advance     int64                   `json:"advance"`
	Robot       *gomirai.Bot            `json:"robot"`
}

func (ticker *JobTicker) RunAtAdvance() { //在提前的时间定时
	timeStamp := time.Now().Unix()
	waitTime := ticker.TargetTime - ticker.Advance*60 - timeStamp
	timer := time.NewTimer(time.Second * time.Duration(waitTime))
	go func() {
		<-timer.C
		ticker.Run()
	}()
}

func (ticker *JobTicker) Run() {
	timer := time.NewTicker(time.Minute * time.Duration(ticker.Interval))
	go func() {
		ticker.Job(ticker)
		for {
			//fmt.Println("ticker running",ticker.Interval,ticker.Advance)
			<-timer.C
			//fmt.Println("ticker perform")
			ticker.Job(ticker)
			timeStamp := time.Now().Unix()
			fmt.Println(timeStamp+ticker.Interval*60, ticker.TargetTime, ticker.TargetTime)
			if timeStamp+ticker.Interval*60 > ticker.TargetTime {
				fmt.Println("ticker at end")
				break
			}
		}
		fmt.Println("a notification has been deleted")
		ticker.FinishJob(ticker)
		//fmt.Println("ticker stop")

	}()
}
