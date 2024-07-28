package model

import (
	"errors"
	"log"
	"math/rand/v2"
	"time"
)

var NonReTriableError = errors.New("non-re-triable error")
var ReTriableError = errors.New("re-triable error")

type JobState string

const (
	Succeeded   JobState = "succeeded"
	Failed      JobState = "failed"
	Progressing JobState = "progressing"
	Unprocessed JobState = "unprocessed"
)

type Result struct {
	Job Job
	Err error
}

type Job struct {
	ID                 int64
	StartAt            time.Time
	ExecutionTime      time.Duration
	State              JobState
	SuccessProbability float64
	Attempts           int
	Results            chan Result
}

// Schedule schedules a job
func (j Job) Schedule() {
	time.AfterFunc(j.StartAt.Sub(time.Now()), j.Execute)
}

func (j Job) Execute() {
	time.Sleep(j.ExecutionTime)
	numb := rand.Float64()

	log.Printf("executed job %#v", j)

	if numb < j.SuccessProbability {
		// report success
		j.Results <- Result{j, nil}
		log.Println("succ reported")
	} else {
		// report fail
		numb = rand.Float64()

		// re-triable happens in 50% of the cases
		if numb >= 0.5 {
			j.Results <- Result{j, NonReTriableError}
		} else {
			j.Results <- Result{j, ReTriableError}
		}
		log.Println("fail reported")
	}

}
