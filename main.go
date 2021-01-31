package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	numOfRoutines, _ := strconv.Atoi(os.Args[1])
	iterCount, _ := strconv.Atoi(os.Args[2])
	sizeArray, _ := strconv.Atoi(os.Args[3])
	fmt.Println("Starting the dispatcher")
	startDispatcher(numOfRoutines, iterCount, sizeArray)
	time.Sleep(10 * time.Second)
}
