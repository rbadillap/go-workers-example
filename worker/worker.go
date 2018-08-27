package worker

import (
	"fmt"
	"github.com/rbadillap/go-workers-example/job"
	"github.com/rbadillap/go-workers-example/task"
	"math/rand"
	"time"
)

var (
	// Define the amount of workers that we are going to use to process all the tasks
	amountOfWorkers = 5
)

func Process(queue task.Task, isComplete chan<- bool) {
	fmt.Println("Working on the task #", queue.Id, " -----")

	// Create two channels for the processes (jobs queue) and the results of each execution
	processes := make(chan job.Job, 20)
	results := make(chan int, 20)

	// Close the channels at the end of the execution
	defer close(processes)
	defer close(results)

	// Send a notification to the queuer channel that we have completed all the tasks
	defer func() {
		isComplete <- true
	}()

	// Based on the amount of workers, execute the process
	for w := 1; w <= amountOfWorkers; w ++ {
		go worker(w, processes, results)
	}

	// According to the amount of jobs of the current task we print the send the goroutine and wait for the result
	for _, currentJob := range queue.Jobs {
		// Send the current job to the worker
		processes <- currentJob

		// Print the result of the execution sent by the worker
		fmt.Println("Result: ", <-results)
	}
}

func worker(id int, processes <-chan job.Job, results chan<- int) {
	for process := range processes {
		// Execute each process and send the result to the channel
		results <- execute(id, process)
	}
}

// This is the function that the worker will "execute" on this case
// we determine the operation to apply to the set of numbers received
func execute(workerId int, j job.Job) int {
	// This part is only for visualization purposes, this can be removed
	duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
	time.Sleep(duration)

	fmt.Println("I am worker #", workerId)
	fmt.Println("Working on the job #", j.Id, " ---")
	fmt.Println("Operation: ", j.Operation, " numbers: ", j.Numbers, " in: ", duration)

	result := 0

	switch j.Operation {
	case "addition":
		result = addition(j.Numbers)
	case "subtract":
		result = subtract(j.Numbers)
	}

	return result
}

func addition(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}

	return total
}

func subtract(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total -= n
	}

	return total
}
