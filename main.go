package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting the dispatcher")
	startDispatcher(5, 5, 100)
	time.Sleep(10 * time.Second)
}
