package jobs

import (
	"errors"
	"github.com/SergeyPanov/job-queue/config"
	"github.com/SergeyPanov/job-queue/db/model"
	"github.com/SergeyPanov/job-queue/db/querier"
	"log"
)

type Scheduler struct {
	config.Config
	JobQuerier *querier.Jobs
	JobQueue   chan model.Job
}

func (js *Scheduler) Schedule() {
	res := make(chan model.Result)
	for {
		select {
		case r := <-res:
			if r.Err == nil {
				log.Println("Job finished successfully")
				r.Job.State = model.Succeeded
				err := js.JobQuerier.Update(r.Job)

				if err != nil {
					log.Println(err)
				}
			} else {
				if errors.Is(r.Err, model.NonReTriableError) {
					r.Job.State = model.Failed
					err := js.JobQuerier.Update(r.Job)

					if err != nil {
						log.Println(err)
					}
				}

				if errors.Is(r.Err, model.ReTriableError) {
					r.Job.State = model.Progressing
					r.Job.Attempts += 1
					err := js.JobQuerier.Update(r.Job)

					if err != nil {
						log.Println(err)
					}

					js.JobQueue <- r.Job
				}
				log.Println("Job finished with error:", r.Err)
			}
		case j := <-js.JobQueue:
			j.Results = res
			log.Printf("Scheduled job %v", j)
			j.Schedule()
		}
	}

}
