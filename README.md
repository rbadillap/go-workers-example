# Workers
PoC about the use of concurrent workers (goroutines) in a "real world" example

## Local Concepts

### Tasks
The tasks are the group (or batch) of jobs that we are going to send to the workers to execute.
Tasks will be sent one-by-one simulating entrypoints as requests received or builds to execute.

### Jobs
Jobs are the process that each existing worker will take to execute.
There could be N jobs per task, and workers should be in charge to take one job to execute and return the result of the operation.

### Worker
The amount of "agents" that we have available to use to execute each job that a task could has.
I am simulating this as the amount of agents that we can assign per client or user in a scenario where we have dynamic amount of agents per user.

Any feedback/suggestion is more than welcome :)

## Example

![Work Example](https://user-images.githubusercontent.com/3170758/44691397-be661480-aa1b-11e8-87be-1ac048e84916.gif "Work example")
