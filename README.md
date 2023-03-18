# Formatter
Distributed sentence formatter implemented in Golang for the Distributed Systems course offered at the University of Tehran.

## sentenceFormatter.go
The formatter implemented in a sequential manner.

It reads an input file line by line and formats each line afterwards.

Finally, it writes the formatted line to a given output file.

## concurrentFormatter.go
The formatter implemented in a concurrent manner.

It reads an input file line by line and creates a Task struct for each line.

The Task struct has an id, which is used for sorting the outputs, and a line string.

The tasks are then sent on a channel for the pool workers to handle i.e. format the lines.

Finally, after each task is done and the line is formatted, the worker sends a TaskResult struct on another channel.

The TaskResult struct also has an id, which is used for sorting the outputs, and a (formatted) line string.

While all the tasks are being done, the coordinator collects the finished tasks from the taskResults channel.

In the end, the coordinator sorts the results and writes them on an output file.

## Notes
The code could be significantly improved regarding both performance and code quality.

This project could be generalised to a map-reduce task. <br>
I recommend checking out [this page](https://pdos.csail.mit.edu/6.824/labs/lab-mr.html) and implement the formatter using the given code structure in the repository.
