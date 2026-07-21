package scheduler

import (
	"context"
	"log"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
)

type FlagExpiredSLAsJob struct {
	VendorService *services.VendorService
}

func (j *FlagExpiredSLAsJob) Name() string {
	return "FlagExpiredSLAsJob"
}

func (j *FlagExpiredSLAsJob) Schedule() string {
	// run every day at 1 AM
	return "0 1 * * *"
}

func (j *FlagExpiredSLAsJob) Execute(ctx context.Context) error {
	log.Println("Running FlagExpiredSLAsJob...")
	return j.VendorService.FlagExpiredSLAs(ctx)
}
