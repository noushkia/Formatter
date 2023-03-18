package phase2

import (
	"bufio"
	"os"
	"sort"
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
	tasks       map[int]*Task
	taskResults map[int]*TaskResult
	tasksLeft   int
}

func (c *Coordinator) Allocate(tasks chan<- *Task) {
	for i := 0; i < len(c.tasks); i++ {
		task := c.tasks[i]
		tasks <- task
	}
	close(tasks)
}

func (c *Coordinator) receiveResults(taskResults <-chan *TaskResult) {
	for taskResult := range taskResults {
		c.tasksLeft--
		c.taskResults[taskResult.id] = taskResult
		if c.tasksLeft == 0 {
			break
		}
	}
}

func (c *Coordinator) HandleResult(output *os.File, taskResults <-chan *TaskResult) error {
	c.receiveResults(taskResults)

	// convert map of task results to a slice
	taskResultSlice := make([]*TaskResult, 0, len(c.taskResults))
	for _, taskResult := range c.taskResults {
		taskResultSlice = append(taskResultSlice, taskResult)
	}

	// sort the slice based on the task id
	sort.Slice(taskResultSlice, func(i, j int) bool {
		return taskResultSlice[i].id < taskResultSlice[j].id
	})

	// write the sorted results to the output file
	for _, taskResult := range taskResultSlice {
		_, err := output.WriteString(taskResult.line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// MakeCoordinator initializes a coordinator with a list of tasks to be processed.
func MakeCoordinator(scanner *bufio.Scanner) (*Coordinator, int, error) {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	c := Coordinator{
		tasks:       make(map[int]*Task),
		taskResults: make(map[int]*TaskResult),
		tasksLeft:   len(lines),
	}

	for i, line := range lines {
		c.tasks[i] = &Task{i, line}
	}

	return &c, c.tasksLeft, nil
}
