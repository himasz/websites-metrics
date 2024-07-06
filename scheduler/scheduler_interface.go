package scheduler

type IScheduler interface {
	AddFunc(spec string, cmd func()) (interface{}, error)
	Start()
}
