package ports

type JobManager interface {
	RegisterJob(job Job) error
	StartScheduler()
}
