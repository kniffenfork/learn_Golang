package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	numOfRoutines, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		fmt.Println("Error parsing number of routines")
	}

	iterCount, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		fmt.Println("Error parsing iteration number")
	}

	sizeArray, err3 := strconv.Atoi(os.Args[3])
	if err3 != nil {
		fmt.Println("Error parsing size of array")
	}

	fmt.Println("Starting the dispatcher")
	startDispatcher(numOfRoutines, iterCount, sizeArray)
	time.Sleep(10 * time.Second)
}
