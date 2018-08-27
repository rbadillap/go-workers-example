package queuer

import (
	"fmt"
	"github.com/rbadillap/go-workers-example/task"
	"github.com/rbadillap/go-workers-example/worker"
)

var (
	queue []task.Task
)

func IsEmpty() bool {
	fmt.Println("Determine whether queue is empty: ", len(queue))
	return len(queue) == 0
}

func Add(task task.Task) {
	fmt.Println("Adding new task to the queue")
	// Adding to the queue
	queue = append(queue, task)

	// Processing the queue
	process()
}

func process() {
	// Create a boolean channel in order to determine if the current queue is complete
	isComplete := make(chan bool, 20)

	// Let the worker to process the queue as a goroutine
	go worker.Process(queue[0], isComplete)

	// If the queue is complete, clear the queue and close the channel
	if <-isComplete {
		close(isComplete)
		queue = make([]task.Task, 0)
	}
}
