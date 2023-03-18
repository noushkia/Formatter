package phase2

import (
	"Formatter/phase1"
	"sync"
)

func Worker(tasks <-chan *Task, taskResults chan<- *TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		taskResult := &TaskResult{id: task.id, line: phase1.Format(task.line)}
		taskResults <- taskResult
	}
}
