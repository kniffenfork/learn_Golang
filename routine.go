package main

type Routine struct {
	ID int
	Quit chan bool
}

func newRoutine(id int) Routine {

	routine := Routine{
		ID: id,
		Quit: make(chan bool),
	}
	return routine
}

func (r *Routine) Start(numOfRoutine int, iterCount int, sizeArr int) {
	go func() {
		for i := 0; i < iterCount; i++ {
			array := GenerateRandomArray(sizeArr)
			Collector(array, numOfRoutine)
		}
	}()
}