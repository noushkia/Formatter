package phase2

import (
	"sync"
)

type Task struct {
	id   int    // task id
	line string // input file line to format
}

type TaskResult struct {
	id   int    // task id
	line string // formatted input file line to write to output file
}

type Coordinator struct {
	mutex sync.Mutex
	tasks map[int]*Task
}

// Done
// find out if the entire job has finished.
//
//func (c *Coordinator) Done() bool {
//	return c.tasksResultsLeft == 0 && c.tasksLeft == 0
//}

func (c *Coordinator) Allocate(tasks chan *Task) error {
	for i := 0; i < len(c.tasks); i++ {
		task := c.tasks[i]
		tasks <- task
	}
	close(tasks)
	return nil
}

func MakeCoordinator(filename string, nWorker int) *Coordinator {
	// todo: get lines from file
	var lines []string

	c := Coordinator{
		tasks: make(map[int]*Task),
	}

	for i, line := range lines {
		c.tasks[i] = &Task{i, line}
	}

	return &c
}
