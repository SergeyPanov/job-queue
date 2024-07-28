package querier

import (
	"database/sql"
	"github.com/SergeyPanov/job-queue/db/model"
	"log"
	"time"
)

type Jobs struct {
	*sql.DB
}

type UpdateJobRequest struct {
	ID       int64
	State    model.JobState
	Attempts int
}

func NewJobs(db *sql.DB) *Jobs {
	return &Jobs{db}
}

func (j *Jobs) LockTx(count int) ([]model.Job, error) {
	tx, err := j.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			log.Fatalf("transaction failed: %v", err)
		} else {
			err = tx.Commit()
			if err != nil {
				log.Fatalf("transaction commit failed: %v", err)
			}
		}
	}()

	query := "SELECT id, start_at, execution_time, state, success_probability, Attempts FROM jobs WHERE Attempts < 3 AND state IN ($1) ORDER BY start_at LIMIT $2"
	rows, err := j.Query(query, model.Unprocessed, count)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	jbs := j.parseJobs(rows)

	for _, jb := range jbs {
		lockQ := "UPDATE jobs SET attempts = attempts + 1, state = $1  WHERE id = $2"
		_, err = j.Exec(lockQ, model.Progressing, jb.ID)
		if err != nil {
			return nil, err
		}
	}

	return jbs, nil
}

func (j *Jobs) Update(job model.Job) error {
	query := `
        UPDATE jobs
        SET state = $1, attempts = $2
        WHERE id = $3
    `
	_, err := j.Exec(query, job.State, job.Attempts, job.ID)
	return err
}

func (j *Jobs) parseJobs(rows *sql.Rows) []model.Job {
	defer rows.Close()
	var jobs []model.Job

	for rows.Next() {
		var job model.Job
		var executionTimeInt int

		err := rows.Scan(&job.ID, &job.StartAt, &executionTimeInt, &job.State, &job.SuccessProbability, &job.Attempts)
		if err != nil {
			log.Println(err)
		}

		job.ExecutionTime = time.Second * time.Duration(executionTimeInt)
		job.Attempts += 1

		jobs = append(jobs, job)
	}

	return jobs
}
