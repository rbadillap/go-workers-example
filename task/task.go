package task

import "github.com/rbadillap/go-workers-example/job"

// Define the Task structure
type Task struct {
	Id   int
	Jobs []job.Job
}
