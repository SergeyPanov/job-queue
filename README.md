# Job Queue

This challenge is about implementing a job queue in Go and Postgres.
The Postgres database should hold the queue of upcoming jobs, and multiple Go instances should function as worker nodes to execute these jobs.


## Requirements

- Multiple worker nodes that execute the jobs that each connect to the database
- Each node can execute multiple jobs in parallel
- When a node goes down during execution of a job, another node should restart the job within 30 seconds
- When a job fails with a "retryable error", it should be retried
- A job may be retried/restarted no more than three times
- Once a job is completed or fails with a "non-retryable error", the job should be marked as "finished" or "failed" in the database
- A single job should only have one active runner working on it
- A job must start with 10 seconds of its scheduled start time

## Dummy job

This challenge is not about implementing a specific job.
However, think about what kind of dummy job you could implement to test your system.

The code of the job does not have to be dynamically fetched from the database and can simply be implemented within the worker node.
