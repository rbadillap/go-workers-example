package main

import (
	"fmt"
	"github.com/rbadillap/go-workers-example/job"
	"github.com/rbadillap/go-workers-example/queuer"
	"github.com/rbadillap/go-workers-example/task"
	"math/rand"
	"time"
)

var (
	// Define the operation that the worker should execute
	operations          = []string{"addition", "subtract"}

	// Define the amount of tasks that we are going to put to the queue
	amountOfTasks       = 5

	// Define the amount of jobs that each task will contain
	amountOfJobsPerTask = rand.Intn(1e1) // random number between 0 and 10
)

func main() {
	// Improve random algorithm with a valid seeder
	rand.Seed(time.Now().Unix())

	fmt.Println("amountOfTasks: ", amountOfTasks)
	fmt.Println("amountOfJobsPerTask: ", amountOfJobsPerTask)

	for t := 1; t <= amountOfTasks; t ++ {
		if queuer.IsEmpty() {
			var jobs []job.Job

			// Create an array of jobs and add them to the queue
			for j := 1; j <= amountOfJobsPerTask; j ++ {
				jobs = append(jobs, job.Job{
					Id:        j,
					Operation: operations[rand.Intn(len(operations))],
					Numbers:   []int{rand.Intn(1e2), rand.Intn(1e3)},
				})
			}

			// Add the new task (with N jobs within) to the queue
			queuer.Add(task.Task{Id: t, Jobs: jobs})
		}
	}
}
