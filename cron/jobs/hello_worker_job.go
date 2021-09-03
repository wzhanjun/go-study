package jobs

import (
	"time"

	"github.com/golang/glog"
)

type helloWorkerJob struct {
}

func (h *helloWorkerJob) Handle() func() {
	return func() {
		now := time.Now().Unix()
		if now%30 == 0 {
			panic("异常")
		}
		glog.Info("hello worker.....")
	}
}

func NewHelloWorkerJob() *helloWorkerJob {
	return &helloWorkerJob{}
}
