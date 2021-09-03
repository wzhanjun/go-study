package jobs

func StartJob() error {

	jobCtx := NewJob()

	if err := jobCtx.RegJob("*/5 * * * * *", NewTestWorkerJob().Build()); err != nil {
		return err
	}

	if err := jobCtx.RegJob("*/5 * * * * *", NewHelloWorkerJob().Handle()); err != nil {
		return err
	}

	go jobCtx.Serve()
	return nil
}
