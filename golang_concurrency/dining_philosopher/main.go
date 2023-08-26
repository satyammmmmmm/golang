package main

import (
	"fmt"
	"sync"
	"time"
)

// philosoper is a struct which store info about philosoper
type Philosoper struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers is list of all philosoper
var philosophers = []Philosoper{
	{"plato", 4, 0},
	{"socrates", 0, 1},
	{"aristotle", 1, 2},
	{"pascal", 2, 3},
	{"locke", 3, 4},
}

// define some variables
var hunger = 3 //how many times does a person eat
var eatTime = 1 * time.Second
var thingTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	//printout welcome message
	fmt.Println("dining Philosopher Problem")
	fmt.Println("**************************")
	fmt.Println("The table is empty")

	//satrt the meal
	dine()
	//printout finish messgae
	fmt.Println("The table is Empty")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	//forks is map of all 5 fork.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}
	//start the meal
	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)

	}
	wg.Wait()
}

func diningProblem(philosoper Philosoper, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("%s is seated at table.\n", philosoper.name)
	seated.Done()
	seated.Wait()

	//eat three times
	for i := hunger; i > 0; i-- {
		if philosoper.leftFork > philosoper.rightFork {
			forks[philosoper.rightFork].Lock()
			fmt.Printf("%s takes the right fork.\n", philosoper.name)
			forks[philosoper.leftFork].Lock()
			fmt.Printf("%s takes the left fork.\n", philosoper.name)

		} else {

			forks[philosoper.leftFork].Lock()
			fmt.Printf("%s takes the left fork.\n", philosoper.name)
			forks[philosoper.rightFork].Lock()
			fmt.Printf("%s takes the right fork.\n", philosoper.name)
		}
		fmt.Printf("%s has both fork.\n", philosoper.name)
		time.Sleep(eatTime)

		fmt.Printf("%s is thinking.\n", philosoper.name)
		time.Sleep(thingTime)
		forks[philosoper.leftFork].Unlock()
		forks[philosoper.rightFork].Unlock()

		fmt.Printf("\t%s put down fork\n,", philosoper.name)

	}

	fmt.Printf(philosoper.name, "is satisfied")
	fmt.Printf(philosoper.name, "left the table")

}
