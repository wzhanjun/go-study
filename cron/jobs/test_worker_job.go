package jobs

import (
	"github.com/golang/glog"
)

type testWorkerJob struct {
}

func (t *testWorkerJob) Build() func() {
	return func() {
		glog.Infof("testWorkerJob执行....")
	}
}

func NewTestWorkerJob() *testWorkerJob {
	return &testWorkerJob{}
}
