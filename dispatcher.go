package main

import (
	"fmt"
	"sort"
)

func startDispatcher(numOfRoutines int, iterCount int, sizeArray int) {
	for i := 0; i < numOfRoutines; i++ {
		fmt.Println("Starting goroutine ", i)
		routine := newRoutine(i)
		routine.Start(i, iterCount, sizeArray)
	}

	go func() {
		for {
			select {
			case task := <-TaskQueue :
					array := task.generatedArray
					sort.Ints(array)
					fmt.Printf( "{%d} {%d} {%d} {%d} {%d}\n", task.routineIdentifier, task.queueInsertionTime, array[0],
										array[len(array) / 2], array[len(array) - 1])
			}
		}
	}()
}
