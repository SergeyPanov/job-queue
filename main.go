package main

import (
	"database/sql"
	"fmt"
	config2 "github.com/SergeyPanov/job-queue/config"
	"github.com/SergeyPanov/job-queue/db/model"
	"github.com/SergeyPanov/job-queue/db/querier"
	"github.com/SergeyPanov/job-queue/jobs"
	"log"
)
import _ "github.com/lib/pq"

func main() {
	c, err := config2.NewConfig("./")
	if err != nil {
		log.Fatalln(err)
	}

	jq := make(chan model.Job, c.MaxJobsToProcess)

	// Database connection string
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s port=%d", c.DbUser, c.DbName, c.DbPassword, c.DbHost, c.DbPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Printf("connected to database: %s\n", connStr)

	jobsQuerier := querier.NewJobs(db)

	scheduler := jobs.Scheduler{
		Config:     c,
		JobQueue:   jq,
		JobQuerier: jobsQuerier,
	}
	go scheduler.Schedule()

	supplier := jobs.NewSupplier(c, jobsQuerier)
	supplier.Supply(jq)
}
