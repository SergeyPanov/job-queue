package jobs

import (
	"github.com/SergeyPanov/job-queue/config"
	"github.com/SergeyPanov/job-queue/db/model"
	"github.com/SergeyPanov/job-queue/db/querier"
	"log"
	"time"
)

type Supplier struct {
	config.Config
	JobsQuerier *querier.Jobs
}

func NewSupplier(cfg config.Config, jq *querier.Jobs) *Supplier {
	return &Supplier{
		Config:      cfg,
		JobsQuerier: jq,
	}
}

func (s *Supplier) Supply(jobsCh chan model.Job) {
	// Query the jobs
	jobsScanTicker := time.NewTicker(time.Duration(s.JobsSupplyRate) * time.Second)
	defer jobsScanTicker.Stop()

	for range jobsScanTicker.C {
		jobs, err := s.JobsQuerier.LockTx(s.MaxJobsToProcess)
		if err != nil {
			log.Printf("error getting jobs: %v", err)
		}

		for _, j := range jobs {
			jobsCh <- j
		}
	}
}
