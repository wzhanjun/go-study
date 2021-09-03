package jobs

import (
	"time"

	"github.com/golang/glog"
	"github.com/robfig/cron/v3"
)

type CronJob interface {
	Serve()
	RegJob(exp string, fn func()) error
}

type jobContext struct {
	*cron.Cron
}

func (s *jobContext) Serve() {
	s.Cron.Start()
}

func (s *jobContext) RegJob(exp string, fn func()) error {
	id, err := s.AddFunc(exp, func() {
		defer func() {
			if p := recover(); p != nil {
				glog.Errorf("执行[%s]定时任务panic, %+v", exp, p)
			}
		}()
		fn()
	})

	if err != nil {
		glog.Errorf("注册定时任务[%s]失败, extryID:%d, err:%+v", exp, id, err)
	}

	return nil
}

func NewJob() CronJob {
	return &jobContext{
		Cron: cron.New(cron.WithSeconds(), cron.WithLocation(time.Local)),
	}
}
