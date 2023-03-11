package phase2

import (
	"Formatter/phase1"
	"sync"
)

func Worker(wg *sync.WaitGroup, tasks chan *Task, taskResults chan *TaskResult) error {
	for task := range tasks {
		output := &TaskResult{id: task.id, line: phase1.Format(task.line)}
		taskResults <- output
	}
	wg.Done()
	return nil
}
