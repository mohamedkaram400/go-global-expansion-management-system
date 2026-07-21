package scheduler

import (
	"context"
	"log"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
	"github.com/robfig/cron/v3"
)

type CronJobManager struct {
	cron *cron.Cron
	jobs []ports.Job
	ctx  context.Context
}

func NewJobManager(ctx context.Context) *CronJobManager {
	return &CronJobManager{
		cron: cron.New(),
		ctx:  ctx,
	}
}

func (m *CronJobManager) RegisterJob(job ports.Job) {
	m.jobs = append(m.jobs, job)
}

func (m *CronJobManager) StartScheduler() {
	for _, job := range m.jobs {

		if _, err := m.cron.AddFunc(job.Schedule(), func() {
			if err := job.Execute(m.ctx); err != nil {
				log.Printf("Error in job %s: %v", job.Name(), err)

			} else {
				log.Printf("Job %s executed successfully", job.Name())
			}
		}); err != nil {
			log.Printf("Failed to schedule job %s: %v", job.Name(), err)
		}

	}

	m.cron.Start()
}
