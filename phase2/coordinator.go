package phase2

import (
	"bufio"
	"log"
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
	tasks     map[int]*Task
	tasksLeft int
}

func (c *Coordinator) Allocate(tasks chan *Task) error {
	for i := 0; i < len(c.tasks); i++ {
		task := c.tasks[i]
		tasks <- task
	}
	return nil
}

func (c *Coordinator) HandleResult(taskResults chan *TaskResult) error {
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(outputFile *os.File, taskResults chan *TaskResult) {
		err := outputFile.Close()
		if err != nil {
			log.Fatal(err)
		}
		close(taskResults)
	}(outputFile, taskResults)

	var formattedLines []*TaskResult
	for taskResult := range taskResults {
		c.tasksLeft--
		formattedLines = append(formattedLines, taskResult)
		if c.tasksLeft == 0 {
			break
		}
	}

	sort.Slice(formattedLines, func(i, j int) bool {
		return formattedLines[i].id < formattedLines[j].id
	})

	for _, taskResult := range formattedLines {
		_, err := outputFile.WriteString(taskResult.line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func MakeCoordinator(filename string) (*Coordinator, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	c := Coordinator{
		tasks:     make(map[int]*Task),
		tasksLeft: len(lines),
	}

	for i, line := range lines {
		c.tasks[i] = &Task{i, line}
	}

	return &c, c.tasksLeft
}
