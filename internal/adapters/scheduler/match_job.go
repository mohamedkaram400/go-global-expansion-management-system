package scheduler

import (
	"context"
	services "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"log"
)

type RefreshMatchesJob struct {
	MatchService   *services.MatchService
	ProjectService *services.ProjectService
}

func (j *RefreshMatchesJob) Name() string {
	return "RefreshMatchesJob"
}

func (j *RefreshMatchesJob) Schedule() string {
	// run every day at midnight
	return "0 * * * *"
}

func (j *RefreshMatchesJob) Execute(ctx context.Context) error {
	log.Println("Running RefreshMatchesJob...")

	// load projects here
	projects, err := j.ProjectService.GetAllActiveProjects(ctx)
	if err != nil {
		return err
	}

	// 2. Loop projects and rebuild matches
	for _, p := range projects {
		_, err := j.MatchService.Rebuild(ctx, p.ID)
		if err != nil {
			log.Printf("failed to rebuild matches for project %d: %v", p.ID, err)
		}
	}

	return nil
}
