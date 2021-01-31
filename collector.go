package main

import "time"

// to collect the tasks for sort the array

type Task struct {
	generatedArray     []int
	routineIdentifier  int
	queueInsertionTime int64
}

var TaskQueue = make(chan Task, 100000)

func Collector(generatedArray []int, routineIdentifier int) {
	currentTask := Task{
		generatedArray:     generatedArray,
		routineIdentifier:  routineIdentifier,
		queueInsertionTime: time.Now().UnixNano() % 10000000,
	}
	TaskQueue <- currentTask
}

